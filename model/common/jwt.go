package common

import (
	"github.com/golang-jwt/jwt/v4"
)

// JwtClaims 自定义jwt的token的数据格式
type JwtClaims struct {
	UserId     uint
	BufferTime int64
	jwt.RegisteredClaims
}
