package response

import "admin-server/model/base"

// SysRoleResponse 角色列表返回实体
type SysRoleResponse struct {
	base.Model
	RoleName *string `json:"roleName"`
	RoleCode *string `json:"roleCode"`
}
