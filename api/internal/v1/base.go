package v1

import (
	"admin-server/model/system/request"
	"admin-server/utils"
	"github.com/gofiber/fiber/v2"
)

// InitBaseRoute 无需验证登录权限的路由
func InitBaseRoute(route *fiber.Router) {
	router := *route
	router.Get("/health", checkHealth)
	router.Post("/login", userLogin)
}

// 健康检查
func checkHealth(context *fiber.Ctx) error {
	return context.JSON(utils.ResponseSuccess("", "ok"))
}

// 用户登录
func userLogin(context *fiber.Ctx) error {
	// 获取登录信息
	var user request.SysUserModel
	_ = context.BodyParser(&user)

	// 校验表单
	if success := utils.ValidateStruct(user); success != nil {
		return context.JSON(&success)
	}

	return context.JSON(userService.UserLogin(context, &user))
}
