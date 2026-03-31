"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { AppShell } from "@/components/AppShell";
import { AuthProvider, useAuth } from "@/lib/auth";

function AuthGuard({ children }: { children: React.ReactNode }) {
  const { isAuthenticated, user } = useAuth();
  const router = useRouter();
  const [isChecking, setIsChecking] = useState(true);

  useEffect(() => {
    // Check auth state after hydration
    const checkAuth = () => {
      const token = localStorage.getItem("bud2_token");
      if (!token) {
        router.push("/login");
      }
      setIsChecking(false);
    };
    
    checkAuth();
  }, [router]);

  // Wait for auth check
  if (isChecking) {
    return (
      <div style={{ 
        minHeight: "100vh", 
        display: "flex", 
        alignItems: "center", 
        justifyContent: "center" 
      }}>
        Carregando...
      </div>
    );
  }

  if (!isAuthenticated) {
    return null;
  }

  return <AppShell user={user}>{children}</AppShell>;
}

export default function AppLayout({ children }: { children: React.ReactNode }) {
  return (
    <AuthProvider>
      <AuthGuard>{children}</AuthGuard>
    </AuthProvider>
  );
}
