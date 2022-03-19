package system

import "admin-server/model/base"

// Role 角色信息
type Role struct {
	base.Model
	Name     *string `gorm:"comment:角色名称;size:100;not null"`
	ParentId *uint   `gorm:"comment:上级ID;not null;default:0;index"`
}
