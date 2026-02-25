package repo

import (
	"context"

	"github.com/brunobotter/feature-flag/application/command"
	"github.com/brunobotter/feature-flag/application/domain"
)

type TenantRepository interface {
	Create(ctx context.Context, cmd command.CreateTenant) (tenant *domain.TenantDomain, err error)
}
