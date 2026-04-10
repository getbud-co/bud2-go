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
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [inviteOpen, setInviteOpen] = useState(false);
  const size = 20;

  const fetchUsers = useCallback(() => {
    setLoading(true);
    usersApi
      .list({
        page,
        size,
      })
      .then(setResult)
      .catch((err: Error) => setError(err.message))
      .finally(() => setLoading(false));
  }, [page]);

  useEffect(() => {
    fetchUsers();
  }, [fetchUsers]);

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
                    Nenhum usuário cadastrado.
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
          await usersApi.create(input);
          fetchUsers();
        }}
      />
    </div>
  );
}
