package role

import (
	"github.com/midaef/emmet-server/app/service"
	"github.com/midaef/emmet-server/config"
)

type RoleEndpoint struct {
	services *service.Services
	config   *config.Config
}

func NewRoleEndpoint(services *service.Services, config *config.Config) *RoleEndpoint {
	return &RoleEndpoint{
		config:   config,
		services: services,
	}
}
