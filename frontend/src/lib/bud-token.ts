import { decodeJwt } from "jose";
import { cookies } from "next/headers";

export const BUD_ACCESS_TOKEN_COOKIE = "bud_access_token";
export const BUD_REFRESH_TOKEN_COOKIE = "bud_refresh_token";
const BUD_SESSION_COOKIE = "bud_auth_session";

export interface BudUser {
  id: string;
  name: string;
  email: string;
  status: string;
  is_system_admin: boolean;
}

export interface BudOrganization {
  id: string;
  name: string;
  domain: string;
  workspace: string;
  status: string;
  membership_role?: string;
  membership_status?: string;
}

export interface BudSession {
  access_token?: string;
  refresh_token?: string;
  token_type?: "Bearer";
  user: BudUser;
  active_organization?: BudOrganization | null;
  organizations: BudOrganization[];
}

export class UnauthorizedError extends Error {
  constructor(message = "Unauthorized") {
    super(message);
    this.name = "UnauthorizedError";
  }
}

function getCookieOptions(maxAge?: number) {
  return {
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    sameSite: "lax" as const,
    path: "/",
    ...(maxAge !== undefined ? { maxAge } : {}),
  };
}

export async function getBudToken(): Promise<string> {
  const cookieStore = await cookies();
  const token = cookieStore.get(BUD_ACCESS_TOKEN_COOKIE)?.value;

  if (!token) {
    const refreshed = await attemptTokenRefresh();
    if (refreshed?.access_token) return refreshed.access_token;
    throw new UnauthorizedError();
  }

  // Proactively refresh if the JWT is expired or within 30 s of expiry.
  try {
    const { exp } = decodeJwt(token);
    if (Date.now() >= ((exp ?? 0) * 1000) - 30_000) {
      const refreshed = await attemptTokenRefresh();
      if (refreshed?.access_token) return refreshed.access_token;
      await clearBudSession();
      throw new UnauthorizedError();
    }
  } catch (e) {
    if (e instanceof UnauthorizedError) throw e;
    // Malformed token — let backend reject it
  }

  return token;
}

export async function saveBudSession(session: BudSession): Promise<void> {
  const cookieStore = await cookies();
  const token = session.access_token;

  if (!token) {
    throw new Error("Missing access token in Bud session");
  }

  // Access token: short-lived, store as session cookie (expires with browser close)
  cookieStore.set(BUD_ACCESS_TOKEN_COOKIE, token, getCookieOptions());

  // Refresh token: 7-day persistent cookie
  if (session.refresh_token) {
    cookieStore.set(
      BUD_REFRESH_TOKEN_COOKIE,
      session.refresh_token,
      getCookieOptions(7 * 24 * 60 * 60),
    );
  }

  cookieStore.set(
    BUD_SESSION_COOKIE,
    JSON.stringify({
      user: session.user,
      active_organization: session.active_organization ?? null,
      organizations: session.organizations,
    } satisfies Omit<BudSession, "access_token" | "refresh_token" | "token_type">),
    getCookieOptions(),
  );
}

export async function clearBudSession(): Promise<void> {
  const cookieStore = await cookies();

  cookieStore.delete(BUD_ACCESS_TOKEN_COOKIE);
  cookieStore.delete(BUD_REFRESH_TOKEN_COOKIE);
  cookieStore.delete(BUD_SESSION_COOKIE);
}

async function attemptTokenRefresh(): Promise<BudSession | null> {
  const cookieStore = await cookies();
  const refreshToken = cookieStore.get(BUD_REFRESH_TOKEN_COOKIE)?.value;

  if (!refreshToken) return null;

  const apiUrl = process.env.BUD_API_URL;
  const response = await fetch(`${apiUrl}/auth/refresh`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ refresh_token: refreshToken }),
  });

  if (!response.ok) return null;

  const session = (await response.json()) as BudSession;
  await saveBudSession(session);
  return session;
}

export async function getBudSession(): Promise<BudSession> {
  const cookieStore = await cookies();
  const rawSession = cookieStore.get(BUD_SESSION_COOKIE)?.value;
  const token = cookieStore.get(BUD_ACCESS_TOKEN_COOKIE)?.value;

  if (!token) {
    // Access token gone (e.g. browser restarted) — try refresh before giving up.
    const refreshed = await attemptTokenRefresh();
    if (refreshed) return refreshed;
    throw new UnauthorizedError();
  }

  if (rawSession) {
    // Check JWT expiry before trusting the cached session.
    // A 30-second buffer ensures we refresh slightly before actual expiry.
    try {
      const { exp } = decodeJwt(token);
      const expiresAt = (exp ?? 0) * 1000;
      if (Date.now() < expiresAt - 30_000) {
        const session = JSON.parse(rawSession) as Omit<
          BudSession,
          "access_token" | "refresh_token" | "token_type"
        >;
        return { ...session, access_token: token, token_type: "Bearer" };
      }
    } catch {
      // Malformed token — fall through to refresh
    }

    // Token is expired or nearly so — refresh proactively.
    const refreshed = await attemptTokenRefresh();
    if (refreshed) return refreshed;
    await clearBudSession();
    throw new UnauthorizedError();
  }

  const apiUrl = process.env.BUD_API_URL;
  const response = await fetch(`${apiUrl}/auth/session`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
    cache: "no-store",
  });

  if (response.status === 401) {
    const refreshed = await attemptTokenRefresh();
    if (refreshed) return refreshed;
    await clearBudSession();
    throw new UnauthorizedError();
  }

  if (!response.ok) {
    throw new Error(`Failed to fetch Bud session: ${response.status}`);
  }

  const session = (await response.json()) as BudSession;
  session.access_token = token;
  session.token_type = "Bearer";
  await saveBudSession(session);
  return session;
}
