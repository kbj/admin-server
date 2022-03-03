package api

import (
	"admin-server/api/internal"
	"admin-server/common/global"
	"admin-server/common/middleware"
	"admin-server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// RegisterRoute 注册系统所有的路由
func RegisterRoute(app *fiber.App) {
	app.Use(cors.New()) // 是否需要允许跨域
	v1Routes := internal.ApiGroupApp.V1

	// 公开的接口不需要登录校验的部分
	{
		publicRouter := app.Group("")
		publicRoutes := v1Routes.PublicApi

		publicRouter.Get("/health", publicRoutes.Health) // 健康状态
		publicRouter.Post("/login", publicRoutes.Login)  // 用户登录
		global.Logger.Info("初始化匿名路由成功！")
	}

	// 用户部分
	{
		userRouter := app.Group("/user", middleware.AuthLogin())
		userRoutes := v1Routes.UserApi

		userRouter.Get("/get", userRoutes.UserInfo)
	}

	register404Routes(app)
	global.Logger.Info("初始化路由完成！")
}

// 注册404路由，兜底用
func register404Routes(app *fiber.App) {
	app.Use("/**", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusNotFound)
		return c.JSON(utils.ResponseNotFound())
	})
}
