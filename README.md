# Feature flag

ÉPICO A — Banco & Migrations (Postgres)
História A1 — “Como dev, quero subir Postgres local via Docker”

Tarefas

 Adicionar serviço postgres no docker-compose.yml

 Configurar volume persistente (não perder dados)

 Expor porta (ex: 5432) apenas se necessário

 Criar usuário, senha, db via env

 Criar healthcheck do Postgres no compose

Aceite

 docker compose up -d postgres sobe sem erro

 psql/DBeaver conecta usando as vars do .env

História A2 — “Como dev, quero ter migrations versionadas para criar o schema”

Tarefas

 Escolher ferramenta (migrate/goose/atlas) e padronizar pasta migrations/

 Criar 0001_init.up.sql (script acima)

 Criar 0001_init.down.sql

 Adicionar comando make migrate-up e make migrate-down

 (Opcional) Container migrate no compose

Aceite

 migrate up cria todas as tabelas

 migrate down 1 remove tudo com sucesso

 Rodar “up” duas vezes não dá erro (idempotência ok)

História A3 — “Como dev, quero garantir integridade do domínio no banco”

Tarefas

 Confirmar constraints (unicidade tenant+key / rule unique)

 Confirmar check constraints (env name / key regex)

 Criar índices principais

 Garantir FK com ON DELETE CASCADE (limpeza consistente)

Aceite

 Não consigo inserir feature.key inválida

 Não consigo criar duas features iguais no mesmo tenant

 Não consigo criar duas rules iguais (feature, env)

História A4 — “Como dev, quero modelar environments padrão por tenant”

Tarefas

 Decidir se envs são:

 Criados por seed manual, ou

 Criados automaticamente no usecase CreateTenant

 Se for automático:

 Implementar transação: cria tenant + 3 envs

 Garantir rollback se qualquer insert falhar

Aceite

 Todo tenant novo já nasce com dev/staging/prod

ÉPICO B — Persistência (Repos) para suportar o CRUD
História B1 — “Como dev, quero repositório de Tenant”

Tarefas

 Interface TenantRepository (Create/List/GetByID)

 Implementação Postgres

 Queries com RETURNING *

 Teste unitário com mock (ou integração com db)

Aceite

 Consigo criar e listar tenants com dados coerentes

História B2 — “Como dev, quero repositório de Feature”

Tarefas

 Interface FeatureRepository (Create/ListByTenant/Update)

 Implementar Postgres

 Tratar erro de unique (tenant_id, key) e retornar erro de domínio

Aceite

 Criar feature dup retorna erro tratável (não 500 genérico)

História B3 — “Como dev, quero repositório de Rules (enabled)”

Tarefas

 Interface FeatureRuleRepository:

 Upsert (feature_id, env_id) → enabled

 ListByTenantAndEnv (pro EvaluateFlags)

 Implementar query de upsert:

 INSERT ... ON CONFLICT ... DO UPDATE

Aceite

 Toggle funciona tanto se regra não existe quanto se já existe

ÉPICO C — Evaluate Flags (coração da fase 1)
História C1 — “Como cliente, quero obter flags de um tenant+env via endpoint /flags”

Tarefas

 Definir contrato do endpoint: GET /flags?tenant=...&env=...

 Validar tenant e env (existem? env pertence ao tenant?)

 Query no banco que retorna key + enabled

 Responder JSON map[string]bool

Aceite

 Endpoint retorna todas as features do tenant com status do env

 Se uma feature não tiver rule, retorna false (default)

História C2 — “Como sistema, quero usar Redis para cachear o resultado de /flags”

Tarefas

 Conectar Redis (client)

 Definir key flags:{tenant}:{env}

 Implementar:

 Cache hit → retorna sem bater no Postgres

 Cache miss → busca Postgres → seta cache com TTL

 Padronizar serialização (JSON)

Aceite

 Segunda chamada de /flags é mais rápida (cache hit)

 TTL funciona (expira e refaz)

História C3 — “Como admin, ao alterar uma flag, quero invalidar o cache”

Tarefas

 No usecase de toggle, deletar flags:{tenant}:{env}

 Garantir tenant/env corretos (evitar deletar cache errado)

 Teste de invalidação (ao menos unit)

Aceite

 Após toggle, /flags reflete mudança imediatamente