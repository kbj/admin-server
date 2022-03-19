package global

import (
	"admin-server/common/enum"
	"errors"
)

// Token部分
var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

// 自定义业务错误
func (e *CustomErrorType) Error() string {
	return e.Msg
}

func NewError(msg string) error {
	err := &CustomErrorType{
		Code: enum.StatusError,
		Msg:  msg,
	}
	return err
}
