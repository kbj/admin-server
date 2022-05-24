package v1

import (
	"admin-server/api/internal/v1/common"
	"admin-server/api/internal/v1/system"
)

type ApiGroup struct {
	BaseApiGroup   common.ApiGroup
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
