import { cookies } from "next/headers";

export const BUD_ACCESS_TOKEN_COOKIE = "bud_access_token";
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

function getCookieOptions() {
  return {
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    sameSite: "lax" as const,
    path: "/",
  };
}

export async function getBudToken(): Promise<string> {
  const cookieStore = await cookies();
  const token = cookieStore.get(BUD_ACCESS_TOKEN_COOKIE)?.value;

  if (!token) {
    throw new UnauthorizedError();
  }

  return token;
}

export async function saveBudSession(session: BudSession): Promise<void> {
  const cookieStore = await cookies();
  const token = session.access_token;

  if (!token) {
    throw new Error("Missing access token in Bud session");
  }

  cookieStore.set(BUD_ACCESS_TOKEN_COOKIE, token, getCookieOptions());
  cookieStore.set(
    BUD_SESSION_COOKIE,
    JSON.stringify({
      user: session.user,
      active_organization: session.active_organization ?? null,
      organizations: session.organizations,
    } satisfies Omit<BudSession, "access_token" | "token_type">),
    getCookieOptions(),
  );
}

export async function clearBudSession(): Promise<void> {
  const cookieStore = await cookies();

  cookieStore.delete(BUD_ACCESS_TOKEN_COOKIE);
  cookieStore.delete(BUD_SESSION_COOKIE);
}

export async function getBudSession(): Promise<BudSession> {
  const cookieStore = await cookies();
  const rawSession = cookieStore.get(BUD_SESSION_COOKIE)?.value;
  const token = cookieStore.get(BUD_ACCESS_TOKEN_COOKIE)?.value;

  if (!token) {
    throw new UnauthorizedError();
  }

  if (rawSession) {
    const session = JSON.parse(rawSession) as Omit<
      BudSession,
      "access_token" | "token_type"
    >;

    return {
      ...session,
      access_token: token,
      token_type: "Bearer",
    };
  }

  const apiUrl = process.env.BUD_API_URL;
  const response = await fetch(`${apiUrl}/auth/session`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
    cache: "no-store",
  });

  if (response.status === 401) {
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
