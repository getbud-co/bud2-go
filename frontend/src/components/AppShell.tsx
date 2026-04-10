"use client";

import { useState } from "react";
import { usePathname } from "next/navigation";
import {
  Avatar,
  Sidebar,
  SidebarHeader,
  SidebarDivider,
  SidebarNav,
  SidebarGroup,
  SidebarItem,
  SidebarFooter,
  SidebarUser,
} from "@mdonangelo/bud-ds";
import { Buildings, Gear, Users, SignOut } from "@phosphor-icons/react";
import { useAuth } from "@/lib/auth";

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

export function AppShell({ children }: { children: React.ReactNode }) {
  const [collapsed, setCollapsed] = useState(false);
  const pathname = usePathname();
  const { logout, organizations, activeOrganization, switchOrganization, user } = useAuth();

  const roleLabel = user?.is_system_admin
    ? "Bud Support"
    : activeOrganization?.membership_role
      ? ROLE_LABEL[activeOrganization.membership_role] ?? activeOrganization.membership_role
      : "Usuário";

  return (
    <div style={{ display: "flex", height: "100vh", overflow: "hidden" }}>
      <Sidebar collapsed={collapsed} onCollapse={() => setCollapsed((v) => !v)}>
        <SidebarHeader collapsedContent={<span style={{ fontWeight: 700, color: "var(--color-orange-500)" }}>B</span>}>
          <span style={{ fontWeight: 700, fontSize: "1.125rem", color: "var(--color-orange-500)" }}>bud</span>
        </SidebarHeader>

        <div style={{ padding: "0 1rem 1rem" }}>
          <label style={{ display: "block", fontSize: "0.75rem", color: "var(--color-text-secondary)", marginBottom: "0.375rem" }}>
            Organização ativa
          </label>
          <select
            value={activeOrganization?.id ?? ""}
            onChange={(e) => void switchOrganization(e.target.value)}
            style={{ width: "100%", padding: "0.625rem 0.75rem", borderRadius: "0.5rem", border: "1px solid var(--color-border)" }}
          >
            {!activeOrganization && <option value="">Selecione uma organização</option>}
            {organizations.map((organization) => (
              <option key={organization.id} value={organization.id}>
                {organization.name} ({organization.workspace})
              </option>
            ))}
          </select>
        </div>

        <SidebarDivider />

        <SidebarNav>
          <SidebarGroup label="Geral">
            <SidebarItem icon={Buildings} label="Organizações" href="/organizations" active={pathname.startsWith("/organizations")} />
          </SidebarGroup>

          <SidebarGroup label="Configurações">
            <SidebarItem icon={Users} label="Usuários" href="/settings/users" active={pathname.startsWith("/settings/users")} />
            <SidebarItem icon={Gear} label="Empresa" href="/settings/organization" active={pathname === "/settings/organization"} />
          </SidebarGroup>
        </SidebarNav>

        <SidebarFooter>
          <SidebarUser
            name={user?.name ?? "Usuário"}
            role={roleLabel}
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

      <main style={{ flex: 1, overflow: "auto", backgroundColor: "var(--color-bg-secondary)" }}>{children}</main>
    </div>
  );
}
