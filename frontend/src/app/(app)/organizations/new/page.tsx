"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Card, CardHeader, CardBody, PageHeader, Button } from "@mdonangelo/bud-ds";
import { ArrowLeft } from "@phosphor-icons/react";
import { OrganizationForm } from "@/components/organizations/OrganizationForm";
import { organizations } from "@/lib/organizations";

export default function NewOrganizationPage() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  return (
    <div style={{ padding: "1.5rem" }}>
      <PageHeader title="Nova organização">
        <Button variant="tertiary" leftIcon={ArrowLeft} onClick={() => router.push("/organizations")}>
          Voltar
        </Button>
      </PageHeader>

      <div style={{ marginTop: "1.5rem", maxWidth: "600px" }}>
        {error && (
          <p style={{ color: "var(--color-error-500)", marginBottom: "1rem" }}>{error}</p>
        )}
        <Card>
          <CardHeader title="Dados da organização" />
          <CardBody>
            <OrganizationForm
              mode="create"
              isLoading={isLoading}
              onSubmit={async ({ name, slug }) => {
                setIsLoading(true);
                setError(null);
                try {
                  await organizations.create({ name, slug });
                  router.push("/organizations");
                } catch (err) {
                  setError(err instanceof Error ? err.message : "Erro ao criar organização");
                } finally {
                  setIsLoading(false);
                }
              }}
            />
          </CardBody>
        </Card>
      </div>
    </div>
  );
}
