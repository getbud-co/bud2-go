"use client";

import { createContext, useContext, useEffect, useState, type ReactNode } from "react";

const API_URL = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";
const TOKEN_KEY = "bud2_token";

export interface AuthUser {
  id: string;
  name: string;
  email: string;
  status: string;
  is_system_admin: boolean;
}

export interface AccessibleOrganization {
  id: string;
  name: string;
  domain: string;
  workspace: string;
  status: "active" | "inactive";
  membership_role?: "admin" | "manager" | "collaborator";
  membership_status?: "invited" | "active" | "inactive";
}

interface AuthPayload {
  access_token?: string;
  token_type?: string;
  user: AuthUser;
  active_organization?: AccessibleOrganization;
  organizations: AccessibleOrganization[];
}

interface AuthContextType {
  token: string | null;
  user: AuthUser | null;
  activeOrganization: AccessibleOrganization | null;
  organizations: AccessibleOrganization[];
  isAuthenticated: boolean;
  isLoading: boolean;
  login: (email: string, password: string) => Promise<void>;
  switchOrganization: (organizationId: string) => Promise<void>;
  refreshSession: () => Promise<void>;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

async function request(path: string, init?: RequestInit, token?: string | null) {
  const response = await fetch(`${API_URL}${path}`, {
    ...init,
    headers: {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...(init?.headers ?? {}),
    },
  });

  if (!response.ok) {
    const contentType = response.headers.get("content-type") ?? "";
    if (contentType.includes("application/json")) {
      const error = await response.json();
      throw new Error(error.detail || error.title || "Request failed");
    }
    throw new Error(response.statusText || "Request failed");
  }

  return response.json();
}

export function AuthProvider({ children }: { children: ReactNode }) {
  const [token, setToken] = useState<string | null>(null);
  const [user, setUser] = useState<AuthUser | null>(null);
  const [activeOrganization, setActiveOrganization] = useState<AccessibleOrganization | null>(null);
  const [organizations, setOrganizations] = useState<AccessibleOrganization[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  function applyPayload(payload: AuthPayload, nextToken?: string | null) {
    if (nextToken !== undefined) {
      setToken(nextToken);
    }
    setUser(payload.user);
    setActiveOrganization(payload.active_organization ?? null);
    setOrganizations(payload.organizations ?? []);
  }

  const refreshSession = async () => {
    const currentToken = localStorage.getItem(TOKEN_KEY);
    if (!currentToken) {
      setToken(null);
      setUser(null);
      setActiveOrganization(null);
      setOrganizations([]);
      return;
    }

    const payload = (await request("/auth/session", { method: "GET" }, currentToken)) as AuthPayload;
    applyPayload(payload, currentToken);
  };

  useEffect(() => {
    const bootstrap = async () => {
      const storedToken = localStorage.getItem(TOKEN_KEY);
      if (!storedToken) {
        setIsLoading(false);
        return;
      }

      try {
        const payload = (await request("/auth/session", { method: "GET" }, storedToken)) as AuthPayload;
        applyPayload(payload, storedToken);
      } catch {
        localStorage.removeItem(TOKEN_KEY);
        setToken(null);
        setUser(null);
        setActiveOrganization(null);
        setOrganizations([]);
      } finally {
        setIsLoading(false);
      }
    };

    void bootstrap();
  }, []);

  const login = async (email: string, password: string) => {
    const payload = (await request("/auth/login", {
      method: "POST",
      body: JSON.stringify({ email, password }),
    })) as AuthPayload;

    localStorage.setItem(TOKEN_KEY, payload.access_token ?? "");
    applyPayload(payload, payload.access_token ?? null);
  };

  const switchOrganization = async (organizationId: string) => {
    if (!token) {
      throw new Error("No active session");
    }

    const payload = (await request(
      "/auth/session",
      {
        method: "PUT",
        body: JSON.stringify({ organization_id: organizationId }),
      },
      token,
    )) as AuthPayload;

    localStorage.setItem(TOKEN_KEY, payload.access_token ?? token);
    applyPayload(payload, payload.access_token ?? token);
  };

  const logout = () => {
    localStorage.removeItem(TOKEN_KEY);
    setToken(null);
    setUser(null);
    setActiveOrganization(null);
    setOrganizations([]);
  };

  return (
    <AuthContext.Provider
      value={{
        token,
        user,
        activeOrganization,
        organizations,
        isAuthenticated: !!token,
        isLoading,
        login,
        switchOrganization,
        refreshSession,
        logout,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}

export function getToken(): string | null {
  if (typeof window === "undefined") {
    return null;
  }
  return localStorage.getItem(TOKEN_KEY);
}
