package internal

import v1 "admin-server/api/internal/v1"

type ApiVersion struct {
	V1 v1.ApiGroup
}

var ApiVersionApp = new(ApiVersion)
