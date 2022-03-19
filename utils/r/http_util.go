package r

import (
	"admin-server/common/enum"
	"admin-server/model/base"
	"github.com/gofiber/fiber/v2"
)

type ResponseOptions func(*base.R)

const (
	defaultOkMsg = "操作成功"
	noAuthMsg    = "您暂时没有访问此资源的权限！"
	notFoundMsg  = "您请求的资源不存在！"
)

// Msg 设置自定义返回Msg的方法
func Msg(msg string) ResponseOptions {
	return func(r *base.R) {
		r.Msg = msg
	}
}

// Code 设置自定义的编码值
func Code(code int) ResponseOptions {
	return func(r *base.R) {
		r.Code = code
	}
}

// Ok 返回成功数据
func Ok(c *fiber.Ctx, data interface{}, options ...ResponseOptions) error {
	c.Status(fiber.StatusOK)
	r := &base.R{
		Code: enum.StatusSuccess,
		Msg:  defaultOkMsg,
		Data: data,
	}

	// 遍历可选参数调用修改
	for _, op := range options {
		op(r)
	}

	return c.JSON(&r)
}

// Fail 返回失败数据
func Fail(c *fiber.Ctx, msg string) error {
	return c.JSON(&base.R{
		Code: enum.StatusError,
		Msg:  msg,
	})
}

// NotAuth 无权限
func NotAuth(c *fiber.Ctx, options ...ResponseOptions) error {
	c.Status(fiber.StatusForbidden)
	r := &base.R{
		Code: enum.StatusForbidden,
		Msg:  noAuthMsg,
	}

	// 遍历可选参数调用修改
	for _, op := range options {
		op(r)
	}

	return c.JSON(&r)
}

// NotFound 找不到
func NotFound(c *fiber.Ctx) error {
	return c.JSON(&base.R{
		Code: enum.StatusNotFound,
		Msg:  notFoundMsg,
	})
}

func Response(c *fiber.Ctx, resp interface{}) error {
	return c.JSON(resp)
}
