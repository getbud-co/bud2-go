import { api } from "./api";

export interface User {
  id: string;
  name: string;
  email: string;
  role: "admin" | "manager" | "collaborator";
  status: "invited" | "active" | "inactive";
  is_system_admin: boolean;
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
  password: string;
  role: string;
}

export interface UpdateUserInput {
  name: string;
  email: string;
  role: string;
  status: string;
}

export const usersApi = {
  list: (params?: { page?: number; size?: number; status?: string }) => {
    const query = new URLSearchParams();
    if (params?.page) query.set("page", String(params.page));
    if (params?.size) query.set("size", String(params.size));
    if (params?.status) query.set("status", params.status);
    const qs = query.toString();
    return api.get<UserListResult>(`/users${qs ? `?${qs}` : ""}`);
  },

  get: (id: string) => api.get<User>(`/users/${id}`),

  create: (input: CreateUserInput) => api.post<User>("/users", input),

  update: (id: string, input: UpdateUserInput) => api.put<User>(`/users/${id}`, input),
};
