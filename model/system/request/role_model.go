package request

import (
	"admin-server/common/types"
	"admin-server/model/common"
)

// SysRoleParamModel 角色列表查询参数
type SysRoleParamModel struct {
	common.PageModel
	RoleName *string             `json:"roleName"`
	RoleCode *string             `json:"roleCode"`
	CreateAt *[2]*types.UnixTime `json:"createAt"`
}
