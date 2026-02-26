package repo

import (
	"context"

	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/domain"
)

type TenantRepository interface {
	Create(ctx context.Context, cmd command.CreateTenant) (tenant *domain.TenantDomain, err error)
	GetById(ctx context.Context, id string) (tenant *domain.TenantDomain, err error)
	GetAll(ctx context.Context, page, limit int) (tenant *domain.TenantPage, err error)
}
