package auth

import (
	"github.com/midaef/emmet-server/app/service"
	"github.com/midaef/emmet-server/config"
)

type AuthEndpoint struct {
	services *service.Services
	config   *config.Config
}

func NewAuthEndpoint(services *service.Services, config *config.Config) *AuthEndpoint {
	return &AuthEndpoint{
		config:   config,
		services: services,
	}
}
