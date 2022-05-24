package middleware

import (
	"admin-server/common/core"
	"admin-server/common/enum"
	"admin-server/model/system"
	"admin-server/utils/r"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// AuthApi 验证是否有API权限的中间件
func AuthApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 获取请求地址和请求方法
		obj := c.Path()
		act := c.Method()

		// 获取登录用户信息
		loginUser := c.Locals(enum.SystemUserInfo).(*system.LoginUser)
		userIdString := strconv.FormatUint(uint64(loginUser.UserId), 10)

		// 使用casbin判断是否有权限
		e := core.GetCasbin()
		ok, _ := e.Enforce(userIdString, obj, act)
		if ok {
			return c.Next()
		}
		return r.NotAuth(c)
	}
}
