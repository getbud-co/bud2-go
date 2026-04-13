import { NextRequest, NextResponse } from "next/server";
import {
  UnauthorizedError,
  getBudSession,
  getBudToken,
  saveBudSession,
  type BudSession,
} from "@/lib/bud-token";

export async function GET() {
  try {
    const session = await getBudSession();
    return NextResponse.json(session);
  } catch (error) {
    if (error instanceof UnauthorizedError) {
      return NextResponse.json({ detail: "Unauthorized" }, { status: 401 });
    }

    return NextResponse.json(
      { detail: "Failed to fetch session" },
      { status: 500 },
    );
  }
}

export async function PUT(request: NextRequest) {
  try {
    const apiUrl = process.env.BUD_API_URL;
    const token = await getBudToken();
    const body = (await request.json()) as { organization_id?: string };

    const response = await fetch(`${apiUrl}/auth/session`, {
      method: "PUT",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        organization_id: body.organization_id,
      }),
    });

    if (response.status === 401) {
      return NextResponse.json({ detail: "Unauthorized" }, { status: 401 });
    }

    if (!response.ok) {
      const error = await response
        .json()
        .catch(() => ({ detail: "Failed to update session" }));
      return NextResponse.json(error, { status: response.status });
    }

    const session = (await response.json()) as BudSession;
    await saveBudSession(session);
    return NextResponse.json(session);
  } catch (error) {
    if (error instanceof UnauthorizedError) {
      return NextResponse.json({ detail: "Unauthorized" }, { status: 401 });
    }

    return NextResponse.json(
      { detail: "Failed to update session" },
      { status: 500 },
    );
  }
}
