package system

import "admin-server/service"

type ApiGroup struct {
	UserApi
	MenuApi
}

var (
	menuService = service.ServiceApp.System.MenuService
)
