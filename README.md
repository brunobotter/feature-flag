# Feature Flag

------------------------------------------------------------------------

# 🟢 Fase 1 --- CRUD mínimo (Admin API) + Repositórios

## 🎯 Objetivo

Conseguir criar tenant, listar, criar feature e ligar/desligar por
ambiente.

------------------------------------------------------------------------

## 📦 Épico B --- Repositórios Postgres
Finalizado
### TenantRepository 

-   Create
-   List
-   GetByID

### EnvironmentRepository

-   ListByTenant
-   GetByName

### FeatureRepository

-   Create
-   ListByTenant
-   GetByKey

### RuleRepository

-   Upsert(feature_id, env_id, enabled)
-   GetByFeatureAndEnv
-   ListEnabledByTenantEnv (usado no /flags)

### ⚠️ Importante (nível sênior)

Mapear erro de UNIQUE do Postgres → erro de domínio (HTTP 409), nunca
500 genérico.

------------------------------------------------------------------------

## 🟡 Épico A4 --- Environments padrão por tenant

Implementar no usecase/CreateTenant:

-   Transação criando:
    -   tenant
    -   dev
    -   staging
    -   prod
-   Rollback automático se qualquer insert falhar

------------------------------------------------------------------------

## 🌐 Endpoints Admin (HTTP)

POST /tenants\
GET /tenants

POST /tenants/:tenantId/features\
GET /tenants/:tenantId/features

PUT /tenants/:tenantId/features/:key\
PUT /tenants/:tenantId/features/:key/environments/:env/toggle

------------------------------------------------------------------------

### ✅ Resultado da Fase 1

Você já tem um "LaunchDarkly mini" administrável via API.

------------------------------------------------------------------------

# 🟡 Fase 2 --- Evaluate Flags (core) + Cache Redis

## 🎯 Objetivo

Entregar o endpoint que os apps clientes vão consumir.

------------------------------------------------------------------------

## 📦 Épico C --- Evaluate

GET /flags?tenant=...&env=...

### Regras

-   Validar tenant existe
-   Validar env pertence ao tenant
-   Retornar map\[string\]bool
-   Default: se não houver rule → false

------------------------------------------------------------------------

## ⚡ Cache Redis

Key padrão:

flags:{tenant}:{env}

-   TTL (30--60s inicialmente)
-   Cache hit → não consulta Postgres
-   Cache miss → consulta Postgres → salva JSON

------------------------------------------------------------------------

## 🔄 Invalidação

Ao executar toggle:

DEL flags:{tenant}:{env}

------------------------------------------------------------------------

### ✅ Resultado da Fase 2

Serviço performático e pronto para uso real.

------------------------------------------------------------------------

# 🔵 Fase 3 --- Frontend Admin (React)

## 🎯 Objetivo

Eliminar dependência de Postman/Insomnia.

------------------------------------------------------------------------

## 🖥 Funcionalidades mínimas

### Tela Tenants

-   Listar
-   Criar

### Tela Features

-   Listar
-   Criar
-   Editar descrição

### Tela Toggle por Environment

-   Grid feature × env
-   Toggle on/off

------------------------------------------------------------------------

## 🧠 Qualidade esperada

-   Client HTTP tipado
-   Estados: loading / error / success
-   Toasts
-   Validação básica

------------------------------------------------------------------------

### ✅ Resultado da Fase 3

Administração completa via interface gráfica.

------------------------------------------------------------------------

# 🟣 Fase 4 --- Realtime (WebSocket / SSE)

## 🎯 Objetivo

Atualizar painel automaticamente quando uma flag mudar.

Canal padrão:

tenant:{id}:env:{env}

### Fluxo

-   Toggle publica evento (Redis PubSub ou broadcast local)
-   Frontend assina e atualiza em tempo real

------------------------------------------------------------------------

### 🔥 Diferencial

Consistência + invalidação + UX moderna.

------------------------------------------------------------------------

# 🟠 Fase 5 --- Rollout gradual e Targeting

## 🎯 Objetivo

Evoluir de booleano simples para sistema completo de feature flags.

------------------------------------------------------------------------

## Evolução do modelo

feature_rules passa a ter:

-   rule_type: BOOLEAN \| PERCENTAGE \| TARGETING
-   percentage: 0--100
-   conditions JSONB

------------------------------------------------------------------------

## Evaluate passa a aceitar:

GET /flags?tenant=...&env=...&userId=123&attributes=... ou POST
/evaluate

------------------------------------------------------------------------

## Regras implementadas

-   Percentage rollout (hash determinístico por userId)
-   Targeting allow/deny
-   Prioridade de regras

------------------------------------------------------------------------

# 🔴 Fase 6 --- Auditoria + Versionamento

## 🎯 Objetivo

Rastrear mudanças e permitir rollback.

-   audit_logs (actor, ação, antes/depois, timestamp)
-   Change request ID
-   Snapshots por feature/env (opcional)

------------------------------------------------------------------------

# 🟤 Fase 7 --- SDK Go

## 🎯 Objetivo

Facilitar adoção nos serviços clientes.

SDK deve ter:

-   Busca automática de flags
-   Cache local
-   Fail-open / Fail-closed
-   Refresh em background
-   Suporte a SSE

------------------------------------------------------------------------

# ⚫ Fase 8 --- Produção (Observabilidade e Segurança)

-   Auth JWT (Admin)
-   RBAC
-   Rate limiting
-   Prometheus
-   OpenTelemetry
-   Logs estruturados
-   CI/CD
-   Migrations automáticas no deploy
