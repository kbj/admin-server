package system

import (
	"admin-server/common/global"
	"admin-server/model/system/request"
	"admin-server/utils/r"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type RoleApi struct{}

// List 角色列表
func (roleApi *RoleApi) List(c *fiber.Ctx) error {
	var param request.SysRoleParamModel
	_ = c.BodyParser(&param)

	err, result := roleService.GetRoleList(&param)
	if err != nil {
		global.Logger.Error("获取角色列表失败", zap.Error(err))
		return r.Fail(c, "获取角色列表失败"+err.Error())
	}
	return r.Ok(c, result)
}
