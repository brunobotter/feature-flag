package viewmodels

import "github.com/brunobotter/feature-flag/application/domain"

type EnviromentViewModel struct {
	Id        string `json:"id"`
	TenantId  string `json:"tenant_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewEnviromentViewModel(e *domain.EnviromentDomain) EnviromentViewModel {
	return EnviromentViewModel{
		Id:        e.Id,
		TenantId:  e.TenantId,
		Name:      e.Name,
		CreatedAt: FormatBRDateTime(e.CreatedAt),
		UpdatedAt: FormatBRDateTime(e.UpdatedAt),
	}
}

func NewEnviromentListViewModel(enviroments []*domain.EnviromentDomain) []EnviromentViewModel {
	list := make([]EnviromentViewModel, 0, len(enviroments))

	for _, enviroment := range enviroments {
		list = append(list, NewEnviromentViewModel(enviroment))
	}

	return list
}
