package enum

import "github.com/gofiber/fiber/v2"

const (
	StatusSuccess             int = 0                               // 请求成功
	StatusForbidden           int = fiber.StatusForbidden           // 无权限
	StatusError               int = -1                              // 请求有错误
	StatusNotFound            int = fiber.StatusNotFound            // 找不到资源
	StatusInternalServerError int = fiber.StatusInternalServerError // 服务器报错
)
