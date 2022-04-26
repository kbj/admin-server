package system

import "admin-server/service"

type ApiGroup struct {
	UserApi
	MenuApi
	RoleApi
}

var (
	menuService = service.ServiceApp.System.MenuService
	roleService = service.ServiceApp.System.RoleService
)
