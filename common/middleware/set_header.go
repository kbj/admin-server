// Package middleware 设置全局都要添加的header
// @author: WeiKai
// @date: 2022/5/22
package middleware

import "github.com/gofiber/fiber/v2"

func SetHeader() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		// 返回header允许读取token的header
		c.Set(fiber.HeaderAccessControlExposeHeaders, fiber.HeaderAuthorization)
		return err
	}
}
