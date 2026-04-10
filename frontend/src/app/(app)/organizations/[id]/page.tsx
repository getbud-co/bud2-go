"use client";

import { useEffect, useState } from "react";
import { useRouter, useParams } from "next/navigation";
import { Card, CardHeader, CardBody, PageHeader, Button } from "@mdonangelo/bud-ds";
import { ArrowLeft } from "@phosphor-icons/react";
import { OrganizationForm } from "@/components/organizations/OrganizationForm";
import { organizations, type Organization } from "@/lib/organizations";

export default function EditOrganizationPage() {
  const router = useRouter();
  const { id } = useParams<{ id: string }>();
  const [org, setOrg] = useState<Organization | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    organizations.get(id).then(setOrg).catch((err: Error) => setError(err.message));
  }, [id]);

  if (!org && !error) {
    return (
      <div style={{ padding: "1.5rem" }}>
        <p>Carregando...</p>
      </div>
    );
  }

  return (
    <div style={{ padding: "1.5rem" }}>
      <PageHeader title={org?.name ?? "Editar organização"}>
        <Button variant="tertiary" leftIcon={ArrowLeft} onClick={() => router.push("/organizations")}>
          Voltar
        </Button>
      </PageHeader>

      <div style={{ marginTop: "1.5rem", maxWidth: "600px" }}>
        {error && (
          <p style={{ color: "var(--color-error-500)", marginBottom: "1rem" }}>{error}</p>
        )}
        {org && (
          <Card>
            <CardHeader title="Dados da organização" />
            <CardBody>
              <OrganizationForm
                mode="edit"
                defaultValues={org}
                isLoading={isLoading}
                onSubmit={async ({ name, domain, workspace, status }) => {
                  setIsLoading(true);
                  setError(null);
                  try {
                    await organizations.update(id, { name, domain, workspace, status });
                    router.push("/organizations");
                  } catch (err) {
                    setError(err instanceof Error ? err.message : "Erro ao salvar organização");
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
