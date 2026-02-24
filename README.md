# Feature flag


Fase 1 — CRUD mínimo (Admin API) + Repositórios

Objetivo: conseguir criar tenant, listar, criar feature, ligar/desligar.

Epico B — Repositórios Postgres

TenantRepository

Create, List, GetByID

EnvironmentRepository

ListByTenant, GetByName

FeatureRepository

Create, ListByTenant, GetByKey

RuleRepository

Upsert(feature_id, env_id, enabled)

GetByFeatureAndEnv

ListEnabledByTenantEnv (pro /flags)

Importante (nível sênior): mapear erro de unique do Postgres → erro de domínio (409), não 500.

Epico A4 — envs padrão por tenant

Implementar no usecase/CreateTenant:

transação: cria tenant + dev/staging/prod

rollback se falhar

Endpoints Admin (HTTP)

POST /tenants

GET /tenants

POST /tenants/:tenantId/features

GET /tenants/:tenantId/features

PUT /tenants/:tenantId/features/:key (editar descrição)

PUT /tenants/:tenantId/features/:key/environments/:env/toggle (enabled true/false)

Resultado da fase 1: você já tem um “LaunchDarkly mini” administrável via API.

Fase 2 — Evaluate Flags (core) + Cache Redis

Objetivo: entregar o endpoint que os apps clientes vão consumir.

Epico C — Evaluate

GET /flags?tenant=...&env=...

valida tenant existe

valida env pertence ao tenant

retorna map[string]bool

default: se não tem rule → false

Cache Redis

key: flags:{tenant}:{env}

TTL (ex: 30s / 60s no começo)

cache miss → busca Postgres → salva JSON

Invalidação

Ao toggle: DEL flags:{tenant}:{env}

Resultado da fase 2: seu serviço já resolve flags rápido e pronto pra uso real.

Fase 3 — Frontend Admin (React)

Objetivo: parar de depender de Postman/Insomnia.

Mínimo do painel

Tela Tenants:

listar + criar

Tela Features do Tenant:

listar + criar + editar descrição

Tela de Toggle por env (dev/staging/prod)

grid: feature x env

toggle on/off

Qualidade

Client HTTP (axios/fetch) com tipagem

Estados: loading/erro/sucesso

Toasts e validação simples

Resultado: você administra tudo por UI.

Fase 4 — Realtime (WebSocket / SSE)

Objetivo: quando uma flag muda, o painel atualiza sem refresh (e opcionalmente os clientes também).

Canal: tenant:{id}:env:{env}

Ao toggle:

publica evento via Redis PubSub (ou só broadcast no processo)

Frontend assina e atualiza a tabela ao vivo

Isso é muito “senior”: consistência, invalidação e UX.

Fase 5 — Rollout gradual e targeting

Objetivo: sair do “enabled bool” e virar feature flag de verdade.

Evoluções na modelagem:

feature_rules vira algo como:

rule_type: BOOLEAN | PERCENTAGE | TARGETING

percentage: 0..100

conditions JSONB (ex: por tenant, por userId, por plano, etc)

Evaluate passa a aceitar:

GET /flags?tenant=...&env=...&userId=123&attributes=...
ou POST /evaluate com body (melhor quando tiver targeting)

Regras:

Percentage rollout (hash determinístico com userId)

Targeting por lista allow/deny

Prioridades (order)

Fase 6 — Auditoria + versionamento

Objetivo: rastrear quem mudou o quê e permitir “voltar”.

audit_logs (actor, ação, antes/depois, timestamp)

“Change request id”

Opcional: snapshots/versionamento por feature/env

Fase 7 — SDK Go (e depois outros)

Objetivo: facilitar adoção.

SDK Go:

busca flags do serviço

cache local

fallback (fail-open/fail-closed)

refresh em background / SSE

Fase 8 — Produção de verdade (observabilidade e segurança)

Auth do Admin (JWT)

RBAC (admin/reader)

Rate limit

Metrics (Prometheus)

Tracing (OpenTelemetry)

Logs estruturados

CI (lint/test/build)

Migrations automatizadas no deploy