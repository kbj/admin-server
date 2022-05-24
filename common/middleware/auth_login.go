package middleware

import (
	"admin-server/common/enum"
	"admin-server/common/global"
	"admin-server/common/global/constants"
	"admin-server/model/common"
	"admin-server/model/system"
	"admin-server/service"
	"admin-server/utils"
	"admin-server/utils/cache"
	"admin-server/utils/r"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"time"
)

var (
	userService = service.ServiceApp.System.UserService
)

// AuthLogin 自定义中间件，校验登录状态
func AuthLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get(fiber.HeaderAuthorization)

		// token是否存在的校验
		if token == "" {
			// token不存在，重新登录
			return r.NeedLogin(c)
		}

		claims, err := utils.ParseJwtToken(token)
		if err != nil || claims == nil {
			// token校验失败
			if err == global.TokenExpired {
				return r.NeedLogin(c)
			}
			return r.NotAuth(c)
		}

		// 获取用户信息保存
		loginUser, err := cache.GetCacheObject[system.LoginUser](constants.REDIS_USER_TOKEN_KEY + token)
		if err != nil || loginUser == nil {
			global.Logger.Warn("从redis获取token对应的登录用户信息失败", zap.Error(err))
			return r.NeedLogin(c)
		} else {
			c.Locals(enum.SystemUserInfo, loginUser)
		}

		// 如果有效期小于5分钟了需要续期token
		if claims.ExpiresAt.Unix()-time.Now().Unix() <= 5*60 {
			newToken, err := renewToken(token, claims, loginUser)
			if err != nil {
				global.Logger.Error("为用户续期登录失败", zap.Error(err))
			} else {
				c.Set(fiber.HeaderAuthorization, newToken)
			}
		}

		return c.Next()
	}
}

// 检查刷新token
func renewToken(oldToken string, claims *common.JwtClaims, loginUser *system.LoginUser) (string, error) {
	claims.ExpiresAt = jwt.NewNumericDate(claims.ExpiresAt.Add(time.Duration(10) * time.Minute))
	newToken, err := utils.RefreshJwtToken(oldToken, claims)
	if err != nil {
		return "", err
	}

	// 更新Redis
	var user system.User
	err = global.Db.First(&user, claims.UserId).Error
	if err != nil {
		return "", err
	}
	loginUser.User = user
	loginUser.ExpireTime = claims.ExpiresAt.Time

	diffTime := claims.ExpiresAt.Time.Unix() - time.Now().Unix()
	err = cache.SetCacheObject(constants.REDIS_USER_TOKEN_KEY+newToken, &loginUser, time.Duration(diffTime)*time.Second)
	if err != nil {
		return "", err
	}
	return newToken, nil
}
