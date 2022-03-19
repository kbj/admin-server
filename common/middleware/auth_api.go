package middleware

import (
	"admin-server/common/core"
	"admin-server/common/enum"
	"admin-server/utils/r"
	"github.com/gofiber/fiber/v2"
)

// AuthApi 验证是否有API权限的中间件
func AuthApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 获取请求地址和请求方法
		obj := c.Path()
		act := c.Method()

		// 获取角色信息
		sub := c.Locals(enum.SystemRoleIds).(*string)

		// 使用casbin判断是否有权限
		e := core.GetCasbin()
		ok, _ := e.Enforce(*sub, obj, act)
		if ok {
			return c.Next()
		}
		return r.NotAuth(c)
	}
}