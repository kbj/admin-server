package boot

import (
	"admin-server/api"
	"admin-server/common/core"
	"admin-server/common/enum"
	"admin-server/common/global"
	"admin-server/model/common"
	"admin-server/utils/r"
	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

// Init 初始化配置
func Init(app *fiber.App) {
	// 初始化配置文件
	global.VP = core.InitializeViper()

	// 初始化Zap日志框架
	global.Logger = core.InitializeZap()

	// 初始化数据库
	global.Db = core.InitializeDbInstance()

	// 初始化Redis连接
	global.RedisClient = core.GetRedisClient()

	// 初始化session池
	global.Session = core.InitializeSession()

	// fiber框架的日志改为zap
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: global.Logger,
	}))

	// 需要使用recover处理错误
	app.Use(recover.New())

	// 注册路由
	api.RegisterRoute(app)
}

// Start 启动服务
func Start(app *fiber.App, listen string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		global.Logger.Info("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(listen); err != nil {
		global.Logger.Error("启动Web服务失败", zap.Error(err))
		os.Exit(0)
	}

	global.Logger.Info("Running cleanup tasks...")
	defer cleanupTasks()
}

// ErrorHandler 通用的错误处理逻辑
func ErrorHandler() func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		httpCode := fiber.StatusInternalServerError
		code := enum.StatusInternalServerError

		if e, ok := err.(*global.CustomErrorType); ok {
			httpCode = fiber.StatusOK
			code = e.Code
		} else if e, ok := err.(*fiber.Error); ok {
			httpCode = e.Code
			if global.Logger != nil {
				global.Logger.Error("Fiber框架发出了异常信息", zap.Error(e))
			}
		} else {
			if global.Logger != nil {
				global.Logger.Error("发生了未知异常", zap.Error(err))
			}
		}

		// 全局使用JSON方式返回错误
		return r.Response(c.Status(httpCode), &common.R{
			Code: code,
			Msg:  err.Error(),
		})
	}
}

// 优雅关机后业务方面需要执行的任务
func cleanupTasks() {
}
