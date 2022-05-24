package system

import (
	"admin-server/model/common"
)

// User 用户信息
type User struct {
	common.Model
	Name     *string `gorm:"index;comment:姓名;size:100;<-:create;unique;not null"`
	RealName *string `gorm:"comment:真实姓名;size:100"`
	Password *string `gorm:"comment:密码:size:1000"`
	Phone    *string `gorm:"index;comment:手机号;size:11;unique;not null"`
}
