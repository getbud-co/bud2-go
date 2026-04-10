"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import {
  Badge,
  Button,
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
import { Plus } from "@phosphor-icons/react";
import { organizations, type Organization, type OrganizationListResult } from "@/lib/organizations";

export default function OrganizationsPage() {
  const router = useRouter();
  const [result, setResult] = useState<OrganizationListResult | null>(null);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const size = 20;

  useEffect(() => {
    setLoading(true);
    organizations
      .list({ page, size })
      .then(setResult)
      .catch((err: Error) => setError(err.message))
      .finally(() => setLoading(false));
  }, [page]);

  return (
    <div style={{ padding: "1.5rem" }}>
      <PageHeader title="Organizações">
        <Button leftIcon={Plus} onClick={() => router.push("/organizations/new")}>
          Nova organização
        </Button>
      </PageHeader>

      <div style={{ marginTop: "1.5rem" }}>
        {error && (
          <p style={{ color: "var(--color-error-500)", marginBottom: "1rem" }}>{error}</p>
        )}

        <Table>
          <TableContent>
            <TableHead>
              <TableRow>
                <TableHeaderCell>Nome</TableHeaderCell>
                <TableHeaderCell>Workspace</TableHeaderCell>
                <TableHeaderCell>Status</TableHeaderCell>
                <TableHeaderCell>Criado em</TableHeaderCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {loading ? (
                <TableRow>
                  <TableCell colSpan={4} style={{ textAlign: "center", padding: "2rem" }}>
                    Carregando...
                  </TableCell>
                </TableRow>
              ) : result?.data.length === 0 ? (
                <TableRow>
                  <TableCell colSpan={4} style={{ textAlign: "center", padding: "2rem" }}>
                    Nenhuma organização cadastrada.
                  </TableCell>
                </TableRow>
              ) : (
                result?.data.map((org: Organization) => (
                  <TableRow
                    key={org.id}
                    onClick={() => router.push(`/organizations/${org.id}`)}
                    style={{ cursor: "pointer" }}
                  >
                    <TableCell>{org.name}</TableCell>
                    <TableCell>{org.workspace}</TableCell>
                    <TableCell>
                      <Badge color={org.status === "active" ? "success" : "neutral"}>
                        {org.status === "active" ? "Ativo" : "Inativo"}
                      </Badge>
                    </TableCell>
                    <TableCell>
                      {new Date(org.created_at).toLocaleDateString("pt-BR")}
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
    </div>
  );
}
