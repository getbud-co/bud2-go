# bud2

SaaS multi-tenant. Monorepo with Go backend and Next.js frontend.

## Structure

```
bud2/
├── compose.yml       # Local dev orchestration (backend + frontend + db)
├── Makefile          # Root targets: dev, build, test, lint, compose-up/down
├── backend/          # Go API
│   ├── cmd/api/      # Entrypoint (main.go)
│   ├── internal/     # Private application code
│   │   ├── domain/       # Entities, value objects, domain interfaces (+ repository interfaces)
│   │   ├── handler/      # HTTP handlers + router (transport layer)
│   │   ├── usecase/      # Application services / use cases (business logic)
│   │   ├── infra/        # Concrete implementations (postgres/, redis/, etc.)
│   │   └── config/       # App configuration
│   ├── pkg/          # Shared/public libraries
│   ├── api/          # OpenAPI spec (openapi.yml — source of truth for API contract)
│   ├── configs/      # Config files (YAML, TOML)
│   ├── migrations/   # Database migrations (golang-migrate)
│   ├── sqlc.yml      # sqlc configuration
│   ├── Dockerfile    # Multi-stage production build
│   └── .golangci.yml # Linter configuration
├── frontend/         # Next.js 15 App
│   ├── Dockerfile    # Multi-stage production build (standalone output)
│   └── src/
│       ├── app/          # App Router (pages, layouts)
│       ├── components/   # React components
│       ├── lib/          # Utilities, API client (api.ts), generated types (types.ts)
│       ├── hooks/        # Custom React hooks
│       └── styles/       # Global styles
└── .github/workflows/ci.yml  # CI pipeline (lint, test, build)
```

## Architecture (Clean Architecture)

Dependency flow — handlers depend on use cases, use cases depend on domain, infra implements domain interfaces:

```
cmd/api/main.go (composition root)
  └─ handler/ → usecase/ → domain/
                              ↑
                           infra/ (implements domain/ interfaces)
```

- **domain/**: Entities, value objects, repository interfaces. Zero external dependencies.
- **usecase/**: Application services that orchestrate business logic. Depends only on domain/.
- **handler/**: HTTP transport. Parses requests, calls use cases, formats responses.
- **infra/**: Concrete implementations of domain interfaces (database, external APIs, etc.).
- **config/**: Loaded at composition root (main.go), injected into dependencies.

## Backend Tooling

- **chi** — HTTP router. Superset of net/http, handlers remain `http.HandlerFunc`.
- **sqlc** — Generates typed Go code from SQL queries. Config in `sqlc.yml`, queries in `internal/infra/sqlc/queries/*.sql`. Run `make sqlc-gen` after changing queries or migrations.
- **golang-migrate** — SQL migrations in `migrations/` as `NNNNNN_desc.up.sql` / `NNNNNN_desc.down.sql`.
- **golang-jwt/jwt** — JWT stateless authentication. Middleware validates token and injects claims into `context.Context`.
- **go-playground/validator** — Input validation on handler request DTOs (format rules only). Business invariants stay in domain/usecase.
- **log/slog** — Structured logging (stdlib). JSON handler in production, text in development.
- **OpenTelemetry** — Traces, metrics, slog bridge. No-op exporter by default; swap via `OTEL_EXPORTER_OTLP_ENDPOINT` without code changes. chi middleware + otelpgx for automatic query instrumentation.

## API Contract

- **OpenAPI 3.1** spec-first in `backend/api/openapi.yml` — source of truth.
- TypeScript types generated from spec: `make api-types` → `frontend/src/lib/types.ts`.
- REST conventions: plural kebab-case resources (`/users`, `/order-items`), semantic HTTP verbs.
- Errors follow RFC 7807 (Problem Details) format, HTTP 422 for validation errors.

## Handler-UseCase Pattern

- Each use case has a single `Execute` method receiving a **Command** (struct) or primitives (≤3 params).
- Use case returns domain entities or primitives, never HTTP DTOs.
- Handler parses HTTP request → validates with go-playground/validator → builds Command → calls Execute → transforms result to Response DTO.
- Handler may orchestrate multiple use cases. Extract to Application Service only when the same orchestration is needed in 2+ handlers.
- Domain entities NEVER have `json:` tags — serialization belongs in handler DTOs.

## Multi-Tenancy (ADR-011)

Strategy: **row-level isolation** — shared schema, `tenant_id UUID NOT NULL` on every business table.

- `tenant_id` extracted from JWT claim once in the handler middleware → injected into `context.Context`.
- Handler extracts via `TenantFromContext(ctx)` → passes as value in Command → usecase → repository params → sqlc `WHERE tenant_id = $1`.
- `TenantID` is a named Value Object in `internal/domain/` — not a raw string or UUID.
- Every business table has `tenant_id` + composite index `(tenant_id, id)`.
- Uniqueness constraints are scoped: `UNIQUE (tenant_id, field)`, never `UNIQUE (field)`.
- RLS (Row-Level Security) enabled on critical tables as a second line of defense.
- Onboarding a new tenant = `INSERT INTO tenants` — never DDL.
- Use cases and domain are unaware of auth/middleware — `TenantID` arrives via Command.

## Conventions

### Backend (Go)
- Follow golang-standards/project-layout
- Business logic lives in `internal/usecase/`, never in handlers
- Domain types and repository interfaces in `internal/domain/`
- Concrete implementations (DB, external services) in `internal/infra/`
- Handlers parse requests, orchestrate use cases, format responses
- SQL queries are explicit (sqlc), never hidden behind an ORM
- Validation of format/structure in handler DTOs; validation of business invariants in domain/usecase
- Tests next to the code they test (`_test.go` suffix)

### Frontend (Next.js)
- App Router with `src/` directory
- TypeScript strict mode
- Path alias `@/*` maps to `src/*`
- Components are functional with TypeScript props
- API client at `src/lib/api.ts` — all backend calls go through it
- Types at `src/lib/types.ts` — generated from OpenAPI spec, never written manually

### General
- Environment variables via `.env` files (never committed), `.env.example` for documentation
- Backend runs on port 8080 by default
- Frontend runs on port 3000 by default
- `make dev` to run both services locally
- `make compose-up` to run via Docker Compose
- `make sqlc-gen` after changing SQL queries or migrations
- `make api-types` after changing the OpenAPI spec
