"use client";

import { FormEvent, useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";

export default function LoginPage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setError(null);
    setIsSubmitting(true);

    try {
      const response = await fetch("/api/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });

      if (!response.ok) {
        const body = (await response.json().catch(() => null)) as
          | { detail?: string; title?: string }
          | null;
        throw new Error(body?.detail ?? body?.title ?? "Falha ao autenticar");
      }

      router.replace(searchParams.get("returnTo") || "/");
      router.refresh();
    } catch (submitError) {
      setError(
        submitError instanceof Error
          ? submitError.message
          : "Falha ao autenticar",
      );
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <div className="flex min-h-dvh items-center justify-center bg-[var(--color-caramel-50)] p-6">
      <div className="w-full max-w-md rounded-3xl border border-[var(--color-caramel-200)] bg-white p-8 shadow-sm">
        <div className="mb-8 space-y-2">
          <p className="text-sm font-medium text-[var(--color-orange-600)]">
            Bud
          </p>
          <h1 className="text-3xl font-semibold text-[var(--color-neutral-900)]">
            Entrar
          </h1>
          <p className="text-sm text-[var(--color-neutral-600)]">
            Use as credenciais do backend para acessar o ambiente.
          </p>
        </div>

        <form className="space-y-4" onSubmit={handleSubmit}>
          <label className="block space-y-2">
            <span className="text-sm font-medium text-[var(--color-neutral-800)]">
              E-mail
            </span>
            <input
              className="w-full rounded-2xl border border-[var(--color-caramel-200)] px-4 py-3 outline-none transition focus:border-[var(--color-caramel-700)]"
              type="email"
              value={email}
              onChange={(event) => setEmail(event.target.value)}
              autoComplete="email"
              required
            />
          </label>

          <label className="block space-y-2">
            <span className="text-sm font-medium text-[var(--color-neutral-800)]">
              Senha
            </span>
            <input
              className="w-full rounded-2xl border border-[var(--color-caramel-200)] px-4 py-3 outline-none transition focus:border-[var(--color-caramel-700)]"
              type="password"
              value={password}
              onChange={(event) => setPassword(event.target.value)}
              autoComplete="current-password"
              required
            />
          </label>

          {error ? (
            <p className="rounded-2xl bg-red-50 px-4 py-3 text-sm text-red-700">
              {error}
            </p>
          ) : null}

          <button
            type="submit"
            disabled={isSubmitting}
            className="w-full rounded-2xl bg-[var(--color-neutral-950)] px-4 py-3 text-sm font-semibold text-white transition hover:opacity-90 disabled:cursor-not-allowed disabled:opacity-60"
          >
            {isSubmitting ? "Entrando..." : "Entrar"}
          </button>
        </form>
      </div>
    </div>
  );
}
