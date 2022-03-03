package internal

import v1 "admin-server/api/internal/v1"

type Api struct {
	V1 v1.ApiGroup
}

var ApiGroupApp = new(Api)
