package system

import (
	"time"
)

// LoginUser 登录用户记录的信息
type LoginUser struct {
	UserId     uint      `json:"userId"`     // 用户id
	DeptId     uint      `json:"deptId"`     // 部门ID
	User       User      `json:"user"`       // 用户信息
	LoginTime  time.Time `json:"loginTime"`  // 登录时间
	ExpireTime time.Time `json:"expireTime"` // 过期时间
}
