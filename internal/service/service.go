package service

import (
	"github.com/midaef/emmet-server/internal/repository"
	"github.com/midaef/emmet-server/pkg/helpers"
)

type Services struct {
}

type Dependencies struct {
	Repository *repository.Repositories
	Hasher     *helpers.Md5
	JWTManager *helpers.JWT
}

func NewServices(deps *Dependencies) *Services {
	return &Services{}
}