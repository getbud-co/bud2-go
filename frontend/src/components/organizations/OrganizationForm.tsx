"use client";

import { useState } from "react";
import { Button, Input, Select } from "@mdonangelo/bud-ds";
import type { SelectOption } from "@mdonangelo/bud-ds";
import type { Organization } from "@/lib/organizations";

const STATUS_OPTIONS: SelectOption[] = [
  { value: "active", label: "Ativo" },
  { value: "inactive", label: "Inativo" },
];

interface OrganizationFormValues {
  name: string;
  domain: string;
  workspace: string;
  status: "active" | "inactive";
}

interface OrganizationFormProps {
  defaultValues?: Partial<Organization>;
  onSubmit: (values: OrganizationFormValues) => Promise<void>;
  isLoading?: boolean;
  mode: "create" | "edit";
}

export function OrganizationForm({ defaultValues, onSubmit, isLoading, mode }: OrganizationFormProps) {
  const [name, setName] = useState(defaultValues?.name ?? "");
  const [domain, setDomain] = useState(defaultValues?.domain ?? "");
  const [workspace, setWorkspace] = useState(defaultValues?.workspace ?? "");
  const [status, setStatus] = useState<"active" | "inactive">(defaultValues?.status ?? "active");

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    await onSubmit({ name, domain, workspace, status });
  }

  return (
    <form onSubmit={handleSubmit} style={{ display: "flex", flexDirection: "column", gap: "1rem" }}>
      <Input label="Nome" value={name} onChange={(e) => setName(e.target.value)} placeholder="Ex: BRQ Digital Solutions" required />
      <Input label="Domínio corporativo" value={domain} onChange={(e) => setDomain(e.target.value)} placeholder="Ex: brq.com" required />
      <Input label="Workspace" value={workspace} onChange={(e) => setWorkspace(e.target.value)} placeholder="Ex: brq" required />
      {mode === "edit" && (
        <Select label="Status" options={STATUS_OPTIONS} value={status} onChange={(val) => setStatus(val as "active" | "inactive")} />
      )}
      <div style={{ display: "flex", justifyContent: "flex-end", gap: "0.5rem", paddingTop: "0.5rem" }}>
        <Button type="submit" loading={isLoading}>
          {mode === "create" ? "Criar organização" : "Salvar alterações"}
        </Button>
      </div>
    </form>
  );
}
