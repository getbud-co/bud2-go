import type { Metadata } from "next";
import "@fontsource/inter";
import "@fontsource/plus-jakarta-sans";
import "@mdonangelo/bud-ds/styles";
import "@/styles/globals.css";

export const metadata: Metadata = {
  title: "bud2",
  description: "bud2 application",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="pt-BR">
      <body>{children}</body>
    </html>
  );
}
