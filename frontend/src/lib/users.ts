import { api } from "./api";

export interface User {
  id: string;
  tenant_id: string;
  name: string;
  email: string;
  role: "admin" | "manager" | "collaborator";
  status: "active" | "inactive";
  created_at: string;
  updated_at: string;
}

export interface UserListResult {
  data: User[];
  total: number;
  page: number;
  size: number;
}

export interface CreateUserInput {
  name: string;
  email: string;
  role: string;
}

export interface UpdateUserInput {
  name: string;
  email: string;
  role: string;
  status: string;
}

const tenantHeader = (tenantId: string) => ({ "X-Tenant-ID": tenantId });

export const usersApi = {
  list: (tenantId: string, params?: { page?: number; size?: number; status?: string; search?: string }) => {
    const query = new URLSearchParams();
    if (params?.page) query.set("page", String(params.page));
    if (params?.size) query.set("size", String(params.size));
    if (params?.status) query.set("status", params.status);
    if (params?.search) query.set("search", params.search);
    const qs = query.toString();
    return api.get<UserListResult>(`/users${qs ? `?${qs}` : ""}`, {
      headers: tenantHeader(tenantId),
    });
  },

  get: (tenantId: string, id: string) =>
    api.get<User>(`/users/${id}`, { headers: tenantHeader(tenantId) }),

  create: (tenantId: string, input: CreateUserInput) =>
    api.post<User>("/users", input, { headers: tenantHeader(tenantId) }),

  update: (tenantId: string, id: string, input: UpdateUserInput) =>
    api.put<User>(`/users/${id}`, input, { headers: tenantHeader(tenantId) }),
};
