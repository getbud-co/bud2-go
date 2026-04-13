import type { Metadata } from "next";
import "./globals.css";
import { cookies } from "next/headers";
import { AppProviders } from "@/providers/AppProviders";
import { AppLayout } from "@/presentation/layout/page";

import "@fontsource/crimson-pro/600.css";
import "@fontsource/plus-jakarta-sans/600.css";
import "@fontsource/plus-jakarta-sans/700.css";
import "@fontsource/inter/400.css";
import "@fontsource/inter/500.css";
import "@fontsource/inter/600.css";

export const metadata: Metadata = {
  title: "Bud",
  description: "Bud frontend",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const cookieStore = await cookies();
  const savedOrgId = cookieStore.get("selectedOrgId")?.value;

  return (
    <html lang="pt-BR">
      <body className="antialiased">
        <AppProviders initialOrgId={savedOrgId}>
          <AppLayout>{children}</AppLayout>
        </AppProviders>
      </body>
    </html>
  );
}
