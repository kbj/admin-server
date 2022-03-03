package system

import (
	"admin-server/model/base"
)

type User struct {
	base.Model
	Name     *string `gorm:"index;comment:姓名;<-:create;unique;not null"`
	RealName *string `gorm:"comment:真实姓名"`
	Password *string `gorm:"comment:密码"`
	Phone    *string `gorm:"index;comment:手机号;size:11;unique;not null"`
	Enable   *uint   `gorm:"comment:状态 1是启用0是禁用;size:1;default:1"`
}
