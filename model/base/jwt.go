package base

import "github.com/golang-jwt/jwt/v4"

// CustomClaims 自定义jwt的token的数据格式
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID    int
	Name  string
	Phone string
}
