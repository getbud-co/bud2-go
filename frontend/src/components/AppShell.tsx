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
import { Buildings, Gear, Users, SignOut } from "@phosphor-icons/react";
import { useAuth } from "@/lib/auth";

interface AppShellProps {
  children: React.ReactNode;
  user?: {
    name: string;
    email: string;
    role: string;
  } | null;
}

function getInitials(name: string) {
  return name
    .split(" ")
    .slice(0, 2)
    .map((w) => w[0])
    .join("")
    .toUpperCase();
}

const ROLE_LABEL: Record<string, string> = {
  admin: "Administrador",
  manager: "Gestor",
  collaborator: "Colaborador",
};

export function AppShell({ children, user }: AppShellProps) {
  const [collapsed, setCollapsed] = useState(false);
  const pathname = usePathname();
  const { logout } = useAuth();

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
              icon={Users}
              label="Usuários"
              href="/settings/users"
              active={pathname.startsWith("/settings/users")}
            />
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
            name={user?.name ?? "Usuário"}
            role={user?.role ? ROLE_LABEL[user.role] ?? user.role : "Usuário"}
            avatar={<Avatar initials={user?.name ? getInitials(user.name) : "U"} size="sm" />}
          />
          <button
            onClick={logout}
            style={{
              display: "flex",
              alignItems: "center",
              gap: "0.5rem",
              padding: "0.5rem",
              marginTop: "0.5rem",
              background: "none",
              border: "none",
              cursor: "pointer",
              color: "var(--color-text-secondary)",
              fontSize: "0.875rem",
            }}
          >
            <SignOut size={16} />
            Sair
          </button>
        </SidebarFooter>
      </Sidebar>

      <main style={{ flex: 1, overflow: "auto", backgroundColor: "var(--color-bg-secondary)" }}>
        {children}
      </main>
    </div>
  );
}
