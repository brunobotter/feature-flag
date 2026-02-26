package repo

import (
	"context"
	"math"

	"github.com/brunobotter/feature-flag/application"
	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/domain"
	apprepo "github.com/brunobotter/feature-flag/application/repo"
	"github.com/brunobotter/feature-flag/infra/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TenantPgRepo struct {
	db  *pgxpool.Pool
	log logger.Logger
}

func NewTenantRepo(db *pgxpool.Pool, log logger.Logger) apprepo.TenantRepository {
	return &TenantPgRepo{
		db:  db,
		log: log,
	}
}

func (r *TenantPgRepo) Create(ctx context.Context, cmd command.CreateTenant) (tenant *domain.TenantDomain, err error) {
	const q = `
    insert into tenants (name)
    values ($1)
    returning id, created_at;
  `
	const createDefaultEnvironmentsQ = `
    insert into environments (tenant_id, name)
    values
      ($1, 'dev'),
      ($1, 'stg'),
      ($1, 'prod');
  `

	tenant = &domain.TenantDomain{Name: cmd.Name}

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, application.Wrap(err)
	}
	defer tx.Rollback(ctx)
	err = tx.QueryRow(ctx, q, tenant.Name).Scan(&tenant.Id, &tenant.CreatedAt)
	if err != nil {
		if friendlyErr := mapDatabaseError(err); friendlyErr != nil {
			return nil, friendlyErr
		}
		return nil, application.Wrap(err)
	}
	if _, err = tx.Exec(ctx, createDefaultEnvironmentsQ, tenant.Id); err != nil {
		if friendlyErr := mapDatabaseError(err); friendlyErr != nil {
			return nil, friendlyErr
		}
		return nil, application.Wrap(err)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, application.Wrap(err)
	}

	return tenant, nil
}

func (r *TenantPgRepo) GetById(ctx context.Context, id string) (*domain.TenantDomain, error) {
	const q = `
    select id, name, created_at, updated_at
    from tenants
	where id = $1;
  `
	tenant := &domain.TenantDomain{}
	err := r.db.QueryRow(ctx, q,
		id,
	).Scan(&tenant.Id, &tenant.Name, &tenant.CreatedAt, &tenant.UpdatedAt)
	if err != nil {
		if friendlyErr := mapDatabaseError(err); friendlyErr != nil {
			return nil, friendlyErr
		}
		return nil, application.Wrap(err)
	}

	return tenant, nil
}

func (r *TenantPgRepo) GetAll(ctx context.Context, page, limit int) (*domain.TenantPage, error) {
	const countQ = `
		select count(1)
		from tenants;
	`
	const q = `
		select id, name, created_at, updated_at
		from tenants
		order by created_at desc
		limit $1 offset $2;
	`

	offset := (page - 1) * limit
	total := 0
	if err := r.db.QueryRow(ctx, countQ).Scan(&total); err != nil {
		if friendlyErr := mapDatabaseError(err); friendlyErr != nil {
			return nil, friendlyErr
		}
		return nil, application.Wrap(err)
	}

	rows, err := r.db.Query(ctx, q, limit, offset)
	if err != nil {
		if friendlyErr := mapDatabaseError(err); friendlyErr != nil {
			return nil, friendlyErr
		}
		return nil, application.Wrap(err)
	}
	defer rows.Close()

	tenants := make([]*domain.TenantDomain, 0)

	for rows.Next() {
		t := &domain.TenantDomain{}
		if err := rows.Scan(&t.Id, &t.Name, &t.CreatedAt, &t.UpdatedAt); err != nil {
			if friendlyErr := mapDatabaseError(err); friendlyErr != nil {
				return nil, friendlyErr
			}
			return nil, application.Wrap(err)
		}
		tenants = append(tenants, t)
	}

	if err := rows.Err(); err != nil {
		if friendlyErr := mapDatabaseError(err); friendlyErr != nil {
			return nil, friendlyErr
		}
		return nil, application.Wrap(err)
	}

	return &domain.TenantPage{
		Items:      tenants,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: int(math.Ceil(float64(total) / float64(limit))),
	}, nil
}
func mapDatabaseError(err error) error {
	return mapPostgresValidationError(err, postgresValidationMap{
		ConstraintMessages: map[string]map[string]string{
			uniqueViolationCode: {
				"ux_tenants_name":       "name já cadastrado",
				"ux_env_tenant_name":    "environment já cadastrado para este tenant",
				"ux_feature_tenant_key": "feature key já cadastrada para este tenant",
				"ux_rule_feature_env":   "regra já cadastrada para esta feature e environment",
			},
			foreignKeyViolationCode: {
				"environments_tenant_id_fkey":       "tenant informado não existe",
				"features_tenant_id_fkey":           "tenant informado não existe",
				"feature_rules_feature_id_fkey":     "feature informada não existe",
				"feature_rules_environment_id_fkey": "environment informado não existe",
			},
			checkViolationCode: {
				"chk_environment_name": "environment deve ser dev, stg ou prod",
				"chk_feature_key":      "feature key inválida",
			},
		},
		DefaultMessages: map[string]string{
			uniqueViolationCode:     "registro já cadastrado",
			foreignKeyViolationCode: "referência inválida",
			checkViolationCode:      "valor inválido",
			notNullViolationCode:    "campo obrigatório não informado",
			invalidTextCode:         "valor inválido para o tipo esperado",
		},
		NoRowsMessage: "registro não encontrado",
	})
}
