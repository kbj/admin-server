package middleware

import (
	"admin-server/common/enum"
	"admin-server/common/global"
	"admin-server/service"
	"admin-server/utils"
	"admin-server/utils/r"
	"github.com/gofiber/fiber/v2"
	"strings"
)

var (
	roleService = service.ServiceApp.System.RoleService
	userService = service.ServiceApp.System.UserService
)

// AuthLogin 自定义中间件，校验登录状态
func AuthLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get(fiber.HeaderAuthorization)

		// token是否存在的校验
		if token == "" || !strings.Contains(token, "Bearer") {
			// token不存在，返回403无权限
			return r.NotAuth(c)
		}
		tokenSplit := strings.Split(token, " ")
		if tokenSplit == nil || len(tokenSplit) != 2 {
			// token不存在，返回403无权限
			return r.NotAuth(c)
		}
		token = tokenSplit[1]

		// 校验token是否正确
		claims, err := utils.ParseJwtToken(token)
		if err != nil || claims == nil {
			if err == global.TokenExpired {
				return r.NotAuth(c, r.Msg("授权已过期"))
			}
			return r.NotAuth(c)
		}

		c.Locals(enum.SystemUserClaims, claims)
		// 查出拥有的角色id列表放入
		roleIds, _ := roleService.GetUserRoles(claims.BaseClaims.ID)
		if roleIds != nil && len(*roleIds) > 0 {
			c.Locals(enum.SystemRoleIds, roleIds)
		}
		// 用户信息
		c.Locals(enum.SystemUserInfo, userService.GetUserInfo(claims.BaseClaims.ID))
		return c.Next()
	}
}
