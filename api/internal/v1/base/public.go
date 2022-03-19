package base

import (
	"admin-server/model/system/request"
	"admin-server/utils"
	"admin-server/utils/r"
	"github.com/gofiber/fiber/v2"
)

type PublicApi struct{}

// Health 检测服务的健康状态
func (p *PublicApi) Health(c *fiber.Ctx) error {
	return r.Ok(c, "ok")
}

// Login 用户登录
func (p *PublicApi) Login(context *fiber.Ctx) error {
	// 获取登录信息
	var user request.SysUserModel
	_ = context.BodyParser(&user)

	// 校验表单
	if err := utils.ValidateStruct(user); err != nil {
		return r.Fail(context, err.Error())
	}

	data, err := userService.UserLogin(&user)
	if err != nil {
		return r.Fail(context, err.Error())
	}
	return r.Ok(context, data, r.Msg("登录成功"))
}
