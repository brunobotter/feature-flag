package requests

import "github.com/brunobotter/feature-flag/api/http"

type TenantRequest struct {
	Request http.HttpRequest
	Name    string `json:"name"`
}
