package viewmodels

import (
	"time"

	"github.com/brunobotter/feature-flag/application/domain"
)

type TenantViewModel struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func NewTenantViewModel(t *domain.TenantDomain) TenantViewModel {
	return TenantViewModel{
		Id:        t.Id,
		Name:      t.Name,
		CreatedAt: FormatBRDateTime(t.CreatedAt),
	}
}

// dd/MM/yyyy HH:mm:ss
func FormatBRDateTime(dt time.Time) string {
	return dt.Format("02/01/2006 15:04:05")
}
