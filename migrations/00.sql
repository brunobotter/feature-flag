-- 0001_init.up.sql
BEGIN;

-- UUID generator
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- =========================
-- TENANTS
-- =========================
CREATE TABLE IF NOT EXISTS tenants (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name        TEXT NOT NULL,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
); 

-- (Opcional) evitar tenants duplicados por nome
CREATE UNIQUE INDEX IF NOT EXISTS ux_tenants_name ON tenants (lower(name));

-- =========================
-- ENVIRONMENTS
-- =========================
CREATE TABLE IF NOT EXISTS environments (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  tenant_id   UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
  name        TEXT NOT NULL,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),

  -- Um tenant não pode ter 2 environments com o mesmo nome
  CONSTRAINT ux_env_tenant_name UNIQUE (tenant_id, name),

  -- Restringe valores válidos (fase 1)
  CONSTRAINT chk_environment_name CHECK (name IN ('dev', 'stg', 'prod'))
);

CREATE INDEX IF NOT EXISTS ix_env_tenant_id ON environments (tenant_id);

-- =========================
-- FEATURES
-- =========================
CREATE TABLE IF NOT EXISTS features (
  id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  tenant_id    UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
  key          TEXT NOT NULL,
  description  TEXT NULL,
  created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at   TIMESTAMPTZ NOT NULL DEFAULT now(),

  -- Uma feature key é única por tenant
  CONSTRAINT ux_feature_tenant_key UNIQUE (tenant_id, key),

  -- Validação simples para evitar keys ruins (ajuste se quiser)
  CONSTRAINT chk_feature_key CHECK (key ~ '^[a-z][a-z0-9_]{1,63}$')
);

CREATE INDEX IF NOT EXISTS ix_features_tenant_id ON features (tenant_id);

-- =========================
-- FEATURE RULES
-- (enabled por environment)
-- =========================
CREATE TABLE IF NOT EXISTS feature_rules (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  feature_id      UUID NOT NULL REFERENCES features(id) ON DELETE CASCADE,
  environment_id  UUID NOT NULL REFERENCES environments(id) ON DELETE CASCADE,
  enabled         BOOLEAN NOT NULL DEFAULT false,
  created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),

  -- Uma regra por (feature, environment)
  CONSTRAINT ux_rule_feature_env UNIQUE (feature_id, environment_id)
);

CREATE INDEX IF NOT EXISTS ix_rules_feature_id ON feature_rules (feature_id);
CREATE INDEX IF NOT EXISTS ix_rules_environment_id ON feature_rules (environment_id);

-- =========================
-- UPDATED_AT trigger (opcional, mas deixa consistente)
-- =========================
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_trigger WHERE tgname = 'tg_tenants_updated_at'
  ) THEN
    CREATE TRIGGER tg_tenants_updated_at
    BEFORE UPDATE ON tenants
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();
  END IF;

  IF NOT EXISTS (
    SELECT 1 FROM pg_trigger WHERE tgname = 'tg_environments_updated_at'
  ) THEN
    CREATE TRIGGER tg_environments_updated_at
    BEFORE UPDATE ON environments
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();
  END IF;

  IF NOT EXISTS (
    SELECT 1 FROM pg_trigger WHERE tgname = 'tg_features_updated_at'
  ) THEN
    CREATE TRIGGER tg_features_updated_at
    BEFORE UPDATE ON features
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();
  END IF;

  IF NOT EXISTS (
    SELECT 1 FROM pg_trigger WHERE tgname = 'tg_feature_rules_updated_at'
  ) THEN
    CREATE TRIGGER tg_feature_rules_updated_at
    BEFORE UPDATE ON feature_rules
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();
  END IF;
END $$;

COMMIT;