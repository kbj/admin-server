package system

import (
	"admin-server/common/global"
	"admin-server/model/base"
	"admin-server/model/system"
	"admin-server/model/system/request"
	"admin-server/utils"
	"go.uber.org/zap"
)

type UserService struct{}

// UserLogin 用户登录
func (u *UserService) UserLogin(param *request.SysUserModel) *base.R {
	var user system.User
	global.Db.Where("name = ?", param.Name).Or("phone = ?", param.Name).Where("enable = ?", 1).Limit(1).Find(&user)
	if user.ID < 1 {
		// 用户不存在
		return utils.ResponseFail("用户名或密码错误")
	}

	// 校验密码
	if utils.Md5Encode(param.Password) != *user.Password {
		// 密码错误
		return utils.ResponseFail("用户名或密码错误")
	}

	// 生成token
	claims := base.BaseClaims{
		ID:    int(user.ID),
		Name:  *user.Name,
		Phone: *user.Phone,
	}
	token, err := utils.CreateJwtToken(utils.CreateClaims(claims))
	if err != nil {
		global.Logger.Error("创建token失败，失败原因：", zap.Error(err))
		return utils.ResponseFail("系统错误！请联系管理员")
	}

	result := map[string]interface{}{
		"id":       user.ID,
		"name":     user.Name,
		"realName": user.RealName,
		"token":    token,
	}
	return utils.ResponseSuccess("登录成功", &result)
}
