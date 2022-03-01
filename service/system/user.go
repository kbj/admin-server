package system

import (
	"admin-server/model/base"
	"admin-server/model/system/request"
	"github.com/gofiber/fiber/v2"
)

type UserService struct{}

// UserLogin 用户登录
func (userService *UserService) UserLogin(ctx *fiber.Ctx, userInfo *request.SysUserModel) *base.ResponseEntity {
	return nil
}
