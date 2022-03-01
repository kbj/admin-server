package middleware

import (
	"admin-server/common/enum"
	"admin-server/common/global"
	"admin-server/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// AuthLogin 自定义中间件，校验登录状态
func AuthLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get(fiber.HeaderAuthorization)

		// token是否存在的校验
		if token == "" || strings.Contains(token, "Bearer") {
			// token不存在，返回403无权限
			notAuth := utils.ResponseNotAuth()
			c.Status(notAuth.Code)
			return c.JSON(notAuth)
		}
		tokenSplit := strings.Split(token, " ")
		if tokenSplit == nil || len(tokenSplit) != 2 {
			// token不存在，返回403无权限
			notAuth := utils.ResponseNotAuth()
			c.Status(notAuth.Code)
			return c.JSON(notAuth)
		}
		token = tokenSplit[1]

		// 校验token是否正确
		claims, err := utils.ParseJwtToken(token)
		if err != nil || claims == nil {
			c.Status(enum.StatusForbidden)
			notAuth := utils.ResponseNotAuth()
			if err == utils.TokenExpired {
				notAuth.Data = "授权已过期"
				return c.JSON(notAuth)
			}
			return c.JSON(notAuth)
		}

		// 将claims放入session中
		session, err := global.Session.Get(c)
		if err == nil && session != nil {
			session.Set("claims", claims)
			// 根据业务需求决定是否要查询出user对象放入session中

			return c.Next()
		} else {
			return c.JSON(utils.ResponseFail("服务器错误，暂时无法获取登录状态"))
		}
	}
}
