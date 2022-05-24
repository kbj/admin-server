package common

import "admin-server/service"

type ApiGroup struct {
	PublicApi
}

var (
	userService = service.ServiceApp.System.UserService
)
