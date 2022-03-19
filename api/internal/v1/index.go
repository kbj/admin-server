package v1

import (
	"admin-server/api/internal/v1/base"
	"admin-server/api/internal/v1/system"
)

type ApiGroup struct {
	BaseApiGroup   base.ApiGroup
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
