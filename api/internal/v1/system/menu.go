package system

import (
	"admin-server/common/enum"
	"admin-server/model/system"
	"admin-server/utils/r"
	"github.com/gofiber/fiber/v2"
)

type MenuApi struct{}

// TreeList 前端树形菜单列表
func (m *MenuApi) TreeList(c *fiber.Ctx) error {
	userId := c.Locals(enum.SystemUserInfo).(*system.User).ID
	treeList, err := menuService.TreeList(userId)
	if err != nil {
		return r.Fail(c, err.Error())
	}
	return r.Ok(c, treeList)
}
