package service

import (
	"context"

	"github.com/midaef/emmet-server/internal/api"
	"github.com/midaef/emmet-server/internal/repository"
	"github.com/midaef/emmet-server/pkg/helpers"
)

type AuthService interface {
	AuthWithCredentials(ctx context.Context, req *api.AuthWithCredentialsRequest) (*api.AuthResponseAccessToken, error)
}

type UserService interface {
	CreateUserByAccessToken(ctx context.Context, req *api.CreateUserByAccessTokenRequest) (*api.CreateUserResponseByAccessToken, error)
}

type Services struct {
	AuthService AuthService
	UserService UserService
}

type Dependencies struct {
	Repository *repository.Repositories
	Hasher     *helpers.Md5
	JWTManager *helpers.JWT
}

func NewServices(deps *Dependencies) *Services {
	authService := NewAuthService(deps.Hasher, deps.JWTManager, deps.Repository.AuthRepository, deps.Repository.TokenRepository)
	userService := NewUserService(deps.Hasher, deps.JWTManager, deps.Repository.UserRepository, deps.Repository.AuthRepository)

	return &Services{
		AuthService: authService,
		UserService: userService,
	}
}