CREATE TABLE users (
    id         UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id  UUID        NOT NULL,
    name       TEXT        NOT NULL,
    email      TEXT        NOT NULL,
    role       TEXT        NOT NULL DEFAULT 'collaborator'
                           CHECK (role IN ('admin', 'manager', 'collaborator')),
    status     TEXT        NOT NULL DEFAULT 'active'
                           CHECK (status IN ('active', 'inactive')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_users_tenant_email ON users (tenant_id, email);
CREATE INDEX idx_users_tenant_status ON users (tenant_id, status);
