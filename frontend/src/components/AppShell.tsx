"use client";

import { useState } from "react";
import { usePathname } from "next/navigation";
import Link from "next/link";
import {
  Avatar,
  Sidebar,
  SidebarHeader,
  SidebarOrgSwitcher,
  SidebarDivider,
  SidebarNav,
  SidebarGroup,
  SidebarItem,
  SidebarFooter,
  SidebarUser,
} from "@mdonangelo/bud-ds";
import { Buildings, Gear } from "@phosphor-icons/react";

interface AppShellProps {
  children: React.ReactNode;
}

export function AppShell({ children }: AppShellProps) {
  const [collapsed, setCollapsed] = useState(false);
  const pathname = usePathname();

  return (
    <div style={{ display: "flex", height: "100vh", overflow: "hidden" }}>
      <Sidebar collapsed={collapsed} onCollapse={() => setCollapsed((v) => !v)}>
        <SidebarHeader collapsedContent={<span style={{ fontWeight: 700, color: "var(--color-orange-500)" }}>B</span>}>
          <span style={{ fontWeight: 700, fontSize: "1.125rem", color: "var(--color-orange-500)" }}>bud</span>
        </SidebarHeader>

        <SidebarOrgSwitcher label="Minha Empresa" />

        <SidebarDivider />

        <SidebarNav>
          <SidebarGroup label="Geral">
            <SidebarItem
              icon={Buildings}
              label="Organizações"
              href="/organizations"
              active={pathname.startsWith("/organizations")}
            />
          </SidebarGroup>

          <SidebarGroup label="Configurações">
            <SidebarItem
              icon={Gear}
              label="Empresa"
              href="/settings/organization"
              active={pathname === "/settings/organization"}
            />
          </SidebarGroup>
        </SidebarNav>

        <SidebarFooter>
          <SidebarUser
            name="Admin"
            role="Administrador"
            avatar={<Avatar initials="AD" size="sm" />}
          />
        </SidebarFooter>
      </Sidebar>

      <main style={{ flex: 1, overflow: "auto", backgroundColor: "var(--color-bg-secondary)" }}>
        {children}
      </main>
    </div>
  );
}
