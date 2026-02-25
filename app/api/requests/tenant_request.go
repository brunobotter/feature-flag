package requests

import "github.com/brunobotter/feature-flag/api/http"

type TenantRequest struct {
	http.HttpRequest
	Name string `json:"name"`
}
