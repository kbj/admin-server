package utils

import (
	"admin-server/common/enum"
	"admin-server/model/base"
)

// ResponseSuccess 返回成功数据
func ResponseSuccess(msg string, data interface{}) *base.R {
	if msg == "" {
		msg = "操作成功"
	}
	return &base.R{
		Code: enum.StatusSuccess,
		Msg:  msg,
		Data: data,
	}
}

// ResponseFail 返回失败数据
func ResponseFail(msg string) *base.R {
	return &base.R{
		Code: enum.StatusError,
		Msg:  msg,
	}
}

// ResponseNotAuth 无权限
func ResponseNotAuth() *base.R {
	return &base.R{
		Code: enum.StatusForbidden,
		Msg:  "您暂时没有访问此资源的权限！",
	}
}

// ResponseNotFound 找不到
func ResponseNotFound() *base.R {
	return &base.R{
		Code: enum.StatusNotFound,
		Msg:  "您请求的资源不存在！",
	}
}
