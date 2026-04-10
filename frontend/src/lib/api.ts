const API_URL = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

type RequestOptions = Omit<RequestInit, "body"> & {
  body?: unknown;
};

export interface ApiProblem {
  type?: string;
  title: string;
  status: number;
  detail?: string;
}

export class ApiError extends Error {
  constructor(public readonly problem: ApiProblem) {
    super(problem.detail ?? problem.title);
    this.name = "ApiError";
  }
}

function getToken(): string | null {
  if (typeof window !== "undefined") {
    return localStorage.getItem("bud2_token");
  }
  return null;
}

async function request<T>(path: string, options: RequestOptions = {}): Promise<T> {
  const { body, headers = {}, ...rest } = options;

  // Build headers with auth token if available
  const requestHeaders: HeadersInit = {
    "Content-Type": "application/json",
    ...(headers as Record<string, string>),
  };

  const token = getToken();
  if (token) {
    (requestHeaders as Record<string, string>)["Authorization"] = `Bearer ${token}`;
  }

  const response = await fetch(`${API_URL}${path}`, {
    headers: requestHeaders,
    body: body ? JSON.stringify(body) : undefined,
    ...rest,
  });

  if (!response.ok) {
    // Handle 401 - redirect to login
    if (response.status === 401) {
      if (typeof window !== "undefined") {
        localStorage.removeItem("bud2_token");
        window.location.href = "/login";
      }
    }

    const contentType = response.headers.get("content-type") ?? "";
    if (contentType.includes("application/json") || contentType.includes("application/problem+json")) {
      const problem = (await response.json()) as ApiProblem;
      throw new ApiError(problem);
    }
    throw new ApiError({ title: response.statusText, status: response.status });
  }

  return response.json() as Promise<T>;
}

export const api = {
  get: <T>(path: string, options?: RequestOptions) =>
    request<T>(path, { ...options, method: "GET" }),

  post: <T>(path: string, body?: unknown, options?: RequestOptions) =>
    request<T>(path, { ...options, method: "POST", body }),

  put: <T>(path: string, body?: unknown, options?: RequestOptions) =>
    request<T>(path, { ...options, method: "PUT", body }),

  delete: <T>(path: string, options?: RequestOptions) =>
    request<T>(path, { ...options, method: "DELETE" }),
};
