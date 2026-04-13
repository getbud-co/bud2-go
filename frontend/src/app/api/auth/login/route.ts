import { NextRequest, NextResponse } from "next/server";
import { saveBudSession, type BudSession } from "@/lib/bud-token";

export async function POST(request: NextRequest) {
  const apiUrl = process.env.BUD_API_URL;
  const body = (await request.json()) as {
    email?: string;
    password?: string;
  };

  const response = await fetch(`${apiUrl}/auth/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email: body.email,
      password: body.password,
    }),
  });

  if (!response.ok) {
    const error = await response
      .json()
      .catch(() => ({ detail: "Falha ao autenticar" }));
    return NextResponse.json(error, { status: response.status });
  }

  const session = (await response.json()) as BudSession;
  await saveBudSession(session);

  return NextResponse.json({
    user: session.user,
    activeOrganization: session.active_organization ?? null,
    organizations: session.organizations,
  });
}
