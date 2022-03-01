package main

import (
	"admin-server/common/boot"
	"admin-server/common/global"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		Prefork:      false,
		ErrorHandler: boot.ErrorHandler(),
	})

	// 初始化
	boot.Init(app)

	// 启动服务
	boot.Start(app, fmt.Sprintf("%s:%s", global.Config.System.Listen, global.Config.System.Port))
}
