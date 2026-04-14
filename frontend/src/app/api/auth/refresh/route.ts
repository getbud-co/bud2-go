import { NextResponse } from "next/server";
import { saveBudSession, clearBudSession, type BudSession } from "@/lib/bud-token";
import { cookies } from "next/headers";

const BUD_REFRESH_TOKEN_COOKIE = "bud_refresh_token";

export async function POST() {
  const cookieStore = await cookies();
  const refreshToken = cookieStore.get(BUD_REFRESH_TOKEN_COOKIE)?.value;

  if (!refreshToken) {
    return NextResponse.json({ detail: "No refresh token" }, { status: 401 });
  }

  const apiUrl = process.env.BUD_API_URL;
  const response = await fetch(`${apiUrl}/auth/refresh`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ refresh_token: refreshToken }),
  });

  if (!response.ok) {
    await clearBudSession();
    return NextResponse.json({ detail: "Unauthorized" }, { status: 401 });
  }

  const session = (await response.json()) as BudSession;
  await saveBudSession(session);

  return NextResponse.json({
    user: session.user,
    active_organization: session.active_organization ?? null,
    organizations: session.organizations,
  });
}
