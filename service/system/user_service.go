package system

import (
	"admin-server/common/global"
	"admin-server/common/global/constants"
	"admin-server/model/system"
	"admin-server/model/system/request"
	"admin-server/model/system/response"
	"admin-server/utils"
	"admin-server/utils/cache"
	"errors"
	"go.uber.org/zap"
	"time"
)

type UserService struct{}

// UserLogin 用户登录
func (u *UserService) UserLogin(param *request.SysUserModel) (*response.LoginResp, *string, error) {
	var user system.User
	global.Db.Where("name = ?", param.Name).Or("phone = ?", param.Name).Limit(1).Find(&user)
	if user.ID == nil || *user.ID < 1 || utils.Md5Encode(param.Password) != *user.Password {
		// 用户或密码不存在
		return nil, nil, errors.New("用户名或密码错误")
	}

	// 通过用户ID生成token
	claims := utils.CreateClaims(user.ID)
	token, err := utils.CreateJwtToken(claims)
	if err != nil {
		global.Logger.Error("创建Token失败，失败原因：", zap.Error(err))
		return nil, nil, errors.New("系统错误！请联系管理员")
	}

	// 保存用户对象进Redis
	result := system.LoginUser{
		UserId:     *user.ID,
		User:       user,
		LoginTime:  time.Now(),
		ExpireTime: claims.ExpiresAt.Time,
	}
	diffTime := claims.ExpiresAt.Time.Unix() - time.Now().Unix()
	err = cache.SetCacheObject(constants.REDIS_USER_TOKEN_KEY+token, &result, time.Duration(diffTime)*time.Second)
	if err != nil {
		global.Logger.Error("保存用户登录信息进Redis失败", zap.Error(err))
	}

	resp := response.LoginResp{
		ID:       user.ID,
		Name:     user.Name,
		RealName: user.RealName,
	}
	return &resp, &token, nil
}

// GetUserInfo 根据用户ID查询出用户信息
func (u *UserService) GetUserInfo(userId int) *system.User {
	var user system.User
	if err := global.Db.First(&user, userId).Error; err != nil {
		panic(err)
	}
	return &user
}
