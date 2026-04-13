import { NextResponse } from "next/server";
import { UnauthorizedError, getBudSession } from "@/lib/bud-token";

export async function GET() {
  try {
    const session = await getBudSession();
    const fullName = session.user.name.trim();
    const initials = fullName
      .split(" ")
      .filter(Boolean)
      .map((part) => part[0] ?? "")
      .slice(0, 2)
      .join("")
      .toUpperCase();

    return NextResponse.json({
      id: session.user.id,
      email: session.user.email,
      fullName,
      initials,
      organizationId: session.active_organization?.id ?? null,
      isSystemAdmin: session.user.is_system_admin,
    });
  } catch (error) {
    if (error instanceof UnauthorizedError) {
      return NextResponse.json({ detail: "Unauthorized" }, { status: 401 });
    }

    return NextResponse.json(
      { detail: "Failed to fetch logged user" },
      { status: 500 },
    );
  }
}
