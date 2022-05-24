package system

import "admin-server/model/common"

// Role 角色信息
type Role struct {
	common.Model
	RoleName *string `gorm:"comment:角色名称;size:100;not null"`
	RoleCode *string `gorm:"comment:角色编码;not null;index"`
}
