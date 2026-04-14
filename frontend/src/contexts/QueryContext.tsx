"use client";

import {
  QueryCache,
  QueryClient,
  QueryClientProvider,
} from "@tanstack/react-query";
import { useState } from "react";

function isUnauthorized(error: unknown): boolean {
  return error instanceof Error && error.message.includes("HTTP 401");
}

function redirectToLogin() {
  const returnTo = encodeURIComponent(
    window.location.pathname + window.location.search,
  );
  window.location.href = `/login?returnTo=${returnTo}`;
}

export default function QueryProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [queryClient] = useState(
    () =>
      new QueryClient({
        queryCache: new QueryCache({
          onError: (error) => {
            if (isUnauthorized(error)) redirectToLogin();
          },
        }),
        defaultOptions: {
          queries: {
            staleTime: 60 * 1000, // 1 minuto
          },
        },
      }),
  );

  return (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
}
