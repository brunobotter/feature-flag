package domain

import "time"

type EnviromentDomain struct {
	Id        string    `json:"id"`
	TenantId  string    `json:"tenant_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
