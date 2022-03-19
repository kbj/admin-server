package system

import "admin-server/model/base"

// RoleMenu 角色拥有的菜单
type RoleMenu struct {
	base.Model
	RoleId *uint `gorm:"comment:角色ID;not null;index"`
	Role   *Role `gorm:"foreignKey:RoleId"`
	MenuId *uint `gorm:"comment:菜单ID;not null;index"`
	Menu   *Menu `gorm:"foreignKey:MenuId"`
}
