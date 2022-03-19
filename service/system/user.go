package system

import (
	"admin-server/common/global"
	"admin-server/model/base"
	"admin-server/model/system"
	"admin-server/model/system/request"
	"admin-server/utils"
	"errors"
	"go.uber.org/zap"
)

type UserService struct{}

// UserLogin 用户登录
func (u *UserService) UserLogin(param *request.SysUserModel) (interface{}, error) {
	var user system.User
	global.Db.Where("name = ?", param.Name).Or("phone = ?", param.Name).Limit(1).Find(&user)
	if user.ID == nil || *user.ID < 1 || utils.Md5Encode(param.Password) != *user.Password {
		// 用户或密码不存在
		return nil, errors.New("用户名或密码错误")
	}

	// 生成token
	claims := base.BaseClaims{
		ID:    int(*user.ID),
		Name:  *user.Name,
		Phone: *user.Phone,
	}
	token, err := utils.CreateJwtToken(utils.CreateClaims(claims))
	if err != nil {
		global.Logger.Error("创建Token失败，失败原因：", zap.Error(err))
		return nil, errors.New("系统错误！请联系管理员")
	}

	result := map[string]interface{}{
		"id":       user.ID,
		"name":     user.Name,
		"realName": user.RealName,
		"token":    token,
	}
	return &result, nil
}

// GetUserInfo 根据用户ID查询出用户信息
func (u *UserService) GetUserInfo(userId int) *system.User {
	var user system.User
	if err := global.Db.First(&user, userId).Error; err != nil {
		panic(err)
	}
	return &user
}
