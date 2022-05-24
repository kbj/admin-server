package utils

import (
	"admin-server/common/global"
	"admin-server/model/common"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// CreateJwtToken 创建token
func CreateJwtToken(claims *common.JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, *claims)
	return token.SignedString([]byte(global.Config.Jwt.SigningKey))
}

// CreateClaims 根据传来的值生成Claims
func CreateClaims(userId *uint) *common.JwtClaims {
	newClaims := common.JwtClaims{
		UserId:     *userId,
		BufferTime: global.Config.Jwt.BufferSecond,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    global.Config.Jwt.Issuer,                                                                         // 签名的发行者
			NotBefore: jwt.NewNumericDate(time.Now().Add(10 * -1 * time.Minute)),                                        // 签发时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Jwt.ExpiresSecond) * time.Second)), // 过期时间
		},
	}
	return &newClaims
}

// ParseJwtToken 解析Token
func ParseJwtToken(tokenString string) (*common.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &common.JwtClaims{}, func(token *jwt.Token) (i any, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, global.TokenInvalid
		}
		return []byte(global.Config.Jwt.SigningKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, global.TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, global.TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, global.TokenNotValidYet
			} else {
				return nil, global.TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*common.JwtClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, global.TokenInvalid
	} else {
		return nil, global.TokenInvalid
	}
}

// RefreshJwtToken 刷新新的token
func RefreshJwtToken(oldToken string, claims *common.JwtClaims) (string, error) {
	// 使用并发控制
	v, err, _ := global.ConcurrencyControl.Do("JWT:"+oldToken, func() (any, error) {
		return CreateJwtToken(claims)
	})
	return v.(string), err
}
