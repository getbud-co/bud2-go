import { api } from "./api";

export interface Organization {
  id: string;
  name: string;
  domain: string;
  workspace: string;
  status: "active" | "inactive";
  created_at: string;
  updated_at: string;
}

export interface OrganizationListResult {
  data: Organization[];
  total: number;
  page: number;
  size: number;
}

export interface CreateOrganizationInput {
  name: string;
  domain: string;
  workspace: string;
}

export interface UpdateOrganizationInput {
  name: string;
  domain: string;
  workspace: string;
  status: "active" | "inactive";
}

export const organizations = {
  list: (params?: { page?: number; size?: number; status?: string }) => {
    const query = new URLSearchParams();
    if (params?.page) query.set("page", String(params.page));
    if (params?.size) query.set("size", String(params.size));
    if (params?.status) query.set("status", params.status);
    const qs = query.toString();
    return api.get<OrganizationListResult>(`/organizations${qs ? `?${qs}` : ""}`);
  },

  get: (id: string) => api.get<Organization>(`/organizations/${id}`),

  create: (input: CreateOrganizationInput) => api.post<Organization>("/organizations", input),

  update: (id: string, input: UpdateOrganizationInput) => api.put<Organization>(`/organizations/${id}`, input),
};
