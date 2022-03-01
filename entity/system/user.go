package system

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"index;comment:姓名;<-:create"`
	Realname  string `gorm:"comment:真实姓名"`
	Password  string `gorm:"comment:密码"`
	Cellphone string `gorm:"comment:手机号;size:11"`
	Enable    uint   `gorm:"comment:状态 1是启用0是禁用;size:1"`
}
