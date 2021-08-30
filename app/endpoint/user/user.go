package user

import (
	"github.com/midaef/emmet-server/app/service"
	"github.com/midaef/emmet-server/config"
)

type UserEndpoint struct {
	services *service.Services
	config   *config.Config
}

func NewUserEndpoint(services *service.Services, config *config.Config) *UserEndpoint {
	return &UserEndpoint{
		config:   config,
		services: services,
	}
}
