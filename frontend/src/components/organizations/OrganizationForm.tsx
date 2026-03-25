"use client";

import { useState, useEffect } from "react";
import { Button, Input, Select } from "@mdonangelo/bud-ds";
import type { SelectOption } from "@mdonangelo/bud-ds";
import type { Organization } from "@/lib/organizations";

const STATUS_OPTIONS: SelectOption[] = [
  { value: "active", label: "Ativo" },
  { value: "inactive", label: "Inativo" },
];

interface OrganizationFormValues {
  name: string;
  slug: string;
  status: "active" | "inactive";
}

interface OrganizationFormProps {
  defaultValues?: Partial<Organization>;
  onSubmit: (values: OrganizationFormValues) => Promise<void>;
  isLoading?: boolean;
  mode: "create" | "edit";
}

function toSlug(name: string): string {
  return name
    .toLowerCase()
    .normalize("NFD")
    .replace(/[\u0300-\u036f]/g, "")
    .replace(/[^a-z0-9\s-]/g, "")
    .trim()
    .replace(/\s+/g, "-");
}

export function OrganizationForm({ defaultValues, onSubmit, isLoading, mode }: OrganizationFormProps) {
  const [name, setName] = useState(defaultValues?.name ?? "");
  const [slug, setSlug] = useState(defaultValues?.slug ?? "");
  const [status, setStatus] = useState<"active" | "inactive">(defaultValues?.status ?? "active");
  const [slugTouched, setSlugTouched] = useState(false);

  useEffect(() => {
    if (!slugTouched && mode === "create") {
      setSlug(toSlug(name));
    }
  }, [name, slugTouched, mode]);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    await onSubmit({ name, slug, status });
  }

  return (
    <form onSubmit={handleSubmit} style={{ display: "flex", flexDirection: "column", gap: "1rem" }}>
      <Input
        label="Nome"
        value={name}
        onChange={(e) => setName(e.target.value)}
        placeholder="Ex: Acme Corporation"
        required
      />
      <Input
        label="Slug"
        value={slug}
        onChange={(e) => {
          setSlugTouched(true);
          setSlug(e.target.value);
        }}
        placeholder="ex: acme-corporation"
        required
      />
      {mode === "edit" && (
        <Select
          label="Status"
          options={STATUS_OPTIONS}
          value={status}
          onChange={(val) => setStatus(val as "active" | "inactive")}
        />
      )}
      <div style={{ display: "flex", justifyContent: "flex-end", gap: "0.5rem", paddingTop: "0.5rem" }}>
        <Button type="submit" loading={isLoading}>
          {mode === "create" ? "Criar organização" : "Salvar alterações"}
        </Button>
      </div>
    </form>
  );
}
