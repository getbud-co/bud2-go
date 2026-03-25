"use client";

import { useEffect, useState, useCallback } from "react";
import {
  Avatar,
  Badge,
  Button,
  Input,
  PageHeader,
  Table,
  TableContent,
  TableHead,
  TableBody,
  TableRow,
  TableHeaderCell,
  TableCell,
  TablePagination,
} from "@mdonangelo/bud-ds";
import { MagnifyingGlass, Plus, Upload } from "@phosphor-icons/react";
import { usersApi, type User, type UserListResult } from "@/lib/users";
import { InviteUserModal } from "@/components/users/InviteUserModal";

// Demo tenant ID — will be replaced by JWT claims
const DEMO_TENANT_ID = "00000000-0000-0000-0000-000000000001";

const ROLE_LABEL: Record<string, string> = {
  admin: "Administrador",
  manager: "Gestor",
  collaborator: "Colaborador",
};

function getInitials(name: string) {
  return name
    .split(" ")
    .slice(0, 2)
    .map((w) => w[0])
    .join("")
    .toUpperCase();
}

export default function UsersPage() {
  const [result, setResult] = useState<UserListResult | null>(null);
  const [page, setPage] = useState(1);
  const [search, setSearch] = useState("");
  const [debouncedSearch, setDebouncedSearch] = useState("");
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [inviteOpen, setInviteOpen] = useState(false);
  const size = 20;

  // Debounce search
  useEffect(() => {
    const t = setTimeout(() => setDebouncedSearch(search), 300);
    return () => clearTimeout(t);
  }, [search]);

  const fetchUsers = useCallback(() => {
    setLoading(true);
    usersApi
      .list(DEMO_TENANT_ID, {
        page,
        size,
        search: debouncedSearch || undefined,
      })
      .then(setResult)
      .catch((err: Error) => setError(err.message))
      .finally(() => setLoading(false));
  }, [page, debouncedSearch]);

  useEffect(() => {
    fetchUsers();
  }, [fetchUsers]);

  // Reset page on search change
  useEffect(() => {
    setPage(1);
  }, [debouncedSearch]);

  return (
    <div style={{ padding: "1.5rem" }}>
      <PageHeader title="Usuários" />

      <div style={{ marginTop: "1.5rem" }}>
        {error && (
          <p style={{ color: "var(--color-error-500)", marginBottom: "1rem" }}>{error}</p>
        )}

        <Table>
          <TableContent>
            <TableHead>
              <TableRow>
                {/* CardHeader row inside table head */}
                <TableHeaderCell colSpan={5}>
                  <div style={{ display: "flex", alignItems: "center", justifyContent: "space-between", gap: "1rem", padding: "0.25rem 0" }}>
                    <div style={{ display: "flex", alignItems: "center", gap: "0.5rem" }}>
                      <span style={{ fontWeight: 600, fontSize: "0.9375rem" }}>Usuários</span>
                      {result && (
                        <Badge color="neutral" size="sm">{result.total}</Badge>
                      )}
                    </div>
                    <div style={{ display: "flex", alignItems: "center", gap: "0.5rem" }}>
                      <div style={{ width: "260px" }}>
                        <Input
                          leftIcon={MagnifyingGlass}
                          placeholder="Buscar por nome ou e-mail..."
                          value={search}
                          onChange={(e) => setSearch(e.target.value)}
                          size="sm"
                        />
                      </div>
                      <Button variant="secondary" leftIcon={Upload} size="sm">
                        Importar usuários
                      </Button>
                      <Button leftIcon={Plus} size="sm" onClick={() => setInviteOpen(true)}>
                        Convidar usuário
                      </Button>
                    </div>
                  </div>
                </TableHeaderCell>
              </TableRow>
              <TableRow>
                <TableHeaderCell>Usuário</TableHeaderCell>
                <TableHeaderCell>Função</TableHeaderCell>
                <TableHeaderCell>Status</TableHeaderCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {loading ? (
                <TableRow>
                  <TableCell colSpan={3} style={{ textAlign: "center", padding: "2rem" }}>
                    Carregando...
                  </TableCell>
                </TableRow>
              ) : result?.data.length === 0 ? (
                <TableRow>
                  <TableCell colSpan={3} style={{ textAlign: "center", padding: "2rem" }}>
                    {debouncedSearch ? `Nenhum usuário encontrado para "${debouncedSearch}".` : "Nenhum usuário cadastrado."}
                  </TableCell>
                </TableRow>
              ) : (
                result?.data.map((user: User) => (
                  <TableRow key={user.id}>
                    <TableCell>
                      <div style={{ display: "flex", alignItems: "center", gap: "0.75rem" }}>
                        <Avatar initials={getInitials(user.name)} size="sm" />
                        <div>
                          <div style={{ fontWeight: 500, fontSize: "0.875rem" }}>{user.name}</div>
                          <div style={{ color: "var(--color-text-secondary)", fontSize: "0.8125rem" }}>{user.email}</div>
                        </div>
                      </div>
                    </TableCell>
                    <TableCell>{ROLE_LABEL[user.role] ?? user.role}</TableCell>
                    <TableCell>
                      <Badge color={user.status === "active" ? "success" : "neutral"} size="sm">
                        {user.status === "active" ? "Ativo" : "Inativo"}
                      </Badge>
                    </TableCell>
                  </TableRow>
                ))
              )}
            </TableBody>
          </TableContent>
          {result && result.total > size && (
            <TablePagination
              currentPage={page}
              totalPages={Math.ceil(result.total / size)}
              onPageChange={setPage}
            />
          )}
        </Table>
      </div>

      <InviteUserModal
        open={inviteOpen}
        onClose={() => setInviteOpen(false)}
        onSubmit={async (input) => {
          await usersApi.create(DEMO_TENANT_ID, input);
          fetchUsers();
        }}
      />
    </div>
  );
}
