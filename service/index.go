package service

import "admin-server/service/system"

type Service struct {
	System system.ServiceGroup
}

var ServiceApp = new(Service)
