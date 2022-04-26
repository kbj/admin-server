package request

import "admin-server/model/base"

// SysRoleParamModel 角色列表查询参数
type SysRoleParamModel struct {
	base.PageModel
	RoleName *string `json:"roleName"`
	RoleCode *string `json:"roleCode"`
}
