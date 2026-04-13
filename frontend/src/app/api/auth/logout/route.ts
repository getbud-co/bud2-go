import { NextResponse } from "next/server";
import { clearBudSession } from "@/lib/bud-token";

export async function POST() {
  await clearBudSession();
  return NextResponse.json({ ok: true });
}
