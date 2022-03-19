package system

import "admin-server/model/base"

// UserRole 用户拥有的角色
type UserRole struct {
	base.Model
	UserId *uint `gorm:"comment:用户ID;not null;index"`
	User   *User `gorm:"foreignKey:UserId"`
	RoleId *uint `gorm:"comment:角色ID;not null;index"`
	Role   *Role `gorm:"foreignKey:RoleId"`
}
