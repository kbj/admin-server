package v1

import (
	"admin-server/model/system/request"
	"admin-server/utils"
	"github.com/gofiber/fiber/v2"
)

type PublicApi struct{}

// Health 检测服务的健康状态
func (p *PublicApi) Health(context *fiber.Ctx) error {
	return context.JSON(utils.ResponseSuccess("", "ok"))
}

// Login 用户登录
func (p *PublicApi) Login(context *fiber.Ctx) error {
	// 获取登录信息
	var user request.SysUserModel
	_ = context.BodyParser(&user)

	// 校验表单
	if success := utils.ValidateStruct(user); success != nil {
		return context.JSON(&success)
	}

	return context.JSON(userService.UserLogin(&user))
}
