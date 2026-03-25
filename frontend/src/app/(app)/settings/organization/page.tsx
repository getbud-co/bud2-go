"use client";

import { useEffect, useState } from "react";
import { Card, CardHeader, CardBody, PageHeader } from "@mdonangelo/bud-ds";
import { OrganizationForm } from "@/components/organizations/OrganizationForm";
import { organizations, type Organization } from "@/lib/organizations";

// Tenant self-service: edita a primeira organização cadastrada
export default function SettingsOrganizationPage() {
  const [org, setOrg] = useState<Organization | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState(false);

  useEffect(() => {
    organizations
      .list({ page: 1, size: 1 })
      .then((r) => r.data[0] && setOrg(r.data[0]))
      .catch((err) => setError(err.message));
  }, []);

  if (!org && !error) {
    return (
      <div style={{ padding: "1.5rem" }}>
        <p>Carregando...</p>
      </div>
    );
  }

  return (
    <div style={{ padding: "1.5rem" }}>
      <PageHeader title="Configurações da empresa" />

      <div style={{ marginTop: "1.5rem", maxWidth: "600px" }}>
        {error && (
          <p style={{ color: "var(--color-error-500)", marginBottom: "1rem" }}>{error}</p>
        )}
        {success && (
          <p style={{ color: "var(--color-success-500)", marginBottom: "1rem" }}>
            Alterações salvas com sucesso.
          </p>
        )}
        {org && (
          <Card>
            <CardHeader title="Dados da organização" />
            <CardBody>
              <OrganizationForm
                mode="edit"
                defaultValues={org}
                isLoading={isLoading}
                onSubmit={async ({ name, slug, status }) => {
                  setIsLoading(true);
                  setError(null);
                  setSuccess(false);
                  try {
                    const updated = await organizations.update(org.id, { name, slug, status });
                    setOrg(updated);
                    setSuccess(true);
                  } catch (err) {
                    setError(err instanceof Error ? err.message : "Erro ao salvar");
                  } finally {
                    setIsLoading(false);
                  }
                }}
              />
            </CardBody>
          </Card>
        )}
      </div>
    </div>
  );
}
