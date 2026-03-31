"use client";

import { useState } from "react";
import {
  Modal,
  ModalHeader,
  ModalBody,
  ModalFooter,
  Button,
  Input,
  Select,
} from "@mdonangelo/bud-ds";
import type { SelectOption } from "@mdonangelo/bud-ds";
import type { CreateUserInput } from "@/lib/users";

const ROLE_OPTIONS: SelectOption[] = [
  { value: "collaborator", label: "Colaborador" },
  { value: "manager", label: "Gestor" },
  { value: "admin", label: "Administrador" },
];

interface InviteUserModalProps {
  open: boolean;
  onClose: () => void;
  onSubmit: (input: CreateUserInput) => Promise<void>;
}

export function InviteUserModal({ open, onClose, onSubmit }: InviteUserModalProps) {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [role, setRole] = useState("collaborator");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  function reset() {
    setName("");
    setEmail("");
    setPassword("");
    setRole("collaborator");
    setError(null);
  }

  function handleClose() {
    reset();
    onClose();
  }

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setLoading(true);
    setError(null);
    try {
      await onSubmit({ name, email, password, role });
      reset();
      onClose();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Erro ao convidar usuário");
    } finally {
      setLoading(false);
    }
  }

  return (
    <Modal open={open} onClose={handleClose}>
      <form onSubmit={handleSubmit}>
        <ModalHeader title="Convidar usuário" onClose={handleClose} />
        <ModalBody>
          <div style={{ display: "flex", flexDirection: "column", gap: "1rem" }}>
            {error && (
              <p style={{ color: "var(--color-error-500)", fontSize: "0.875rem" }}>{error}</p>
            )}
            <Input
              label="Nome"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="Ex: Ana Ferreira"
              required
            />
            <Input
              label="E-mail"
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              placeholder="ana@empresa.com"
              required
            />
            <Input
              label="Senha"
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="Mínimo 8 caracteres"
              required
            />
            <Select
              label="Função"
              options={ROLE_OPTIONS}
              value={role}
              onChange={(val) => setRole(val)}
            />
          </div>
        </ModalBody>
        <ModalFooter>
          <Button type="button" variant="tertiary" onClick={handleClose}>
            Cancelar
          </Button>
          <Button type="submit" loading={loading}>
            Convidar
          </Button>
        </ModalFooter>
      </form>
    </Modal>
  );
}
