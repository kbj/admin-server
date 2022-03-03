package v1

import "admin-server/service"

type ApiGroup struct {
	UserApi
	PublicApi
}

var (
	userService = service.ServiceApp.System.UserService
)
