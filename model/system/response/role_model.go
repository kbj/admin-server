package response

import (
	"admin-server/model/common"
)

// SysRoleResponse 角色列表返回实体
type SysRoleResponse struct {
	common.Model
	RoleName *string `json:"roleName"`
	RoleCode *string `json:"roleCode"`
}
