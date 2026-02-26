package requests

import "github.com/brunobotter/feature-flag/api/http"

type CreateTenantRequest struct {
	http.HttpRequest
	Name string `json:"name"`
}

type TenantRequest struct {
	http.HttpRequest
	Name string `json:"name"`
}
