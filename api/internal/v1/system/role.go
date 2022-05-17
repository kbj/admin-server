package system

import (
	"admin-server/common/global"
	"admin-server/model/system/request"
	"admin-server/utils"
	"admin-server/utils/r"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strings"
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

// Delete 删除角色
func (roleApi *RoleApi) Delete(c *fiber.Ctx) error {
	id := strings.Split(c.Params("id"), ",")
	// 将id转为int格式
	intId, err := utils.StringArray2intArray(&id)
	if err != nil {
		return r.Fail(c, "删除失败"+err.Error())
	}

	// 删除角色
	err = roleService.Delete(intId)
	if err != nil {
		if e, ok := err.(*global.CustomErrorType); ok {
			return r.Response(c, e)
		}
		global.Logger.Error("删除角色失败", zap.Error(err))
		return r.Fail(c, "删除失败"+err.Error())
	}
	return r.Ok(c, "删除成功")
}
