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
	DeleteUserByAccessToken(ctx context.Context, req *api.DeleteUserByAccessTokenRequest) (*api.DeleteUserResponseByAccessToken, error)
	UpdatePasswordByAccessToken(ctx context.Context, req *api.UpdatePasswordByAccessTokenRequest) (*api.UpdatePasswordResponseByAccessToken, error)
}

type RoleService interface {
	CreateRoleByAccessToken(ctx context.Context, req *api.CreateRoleByAccessTokenRequest) (*api.CreateRoleResponseByAccessToken, error)
	DeleteRoleByAccessToken(ctx context.Context, req *api.DeleteRoleByAccessTokenRequest) (*api.DeleteRoleResponseByAccessToken, error)
}

type DataService interface {
	CreateValueByAccessToken(ctx context.Context, req *api.CreateValueByAccessTokenRequest) (*api.CreateValueResponseByAccessToken, error)
	DeleteValueByAccessToken(ctx context.Context, req *api.DeleteValueByAccessTokenRequest) (*api.DeleteValueResponseByAccessToken, error)
	GetValueByAccessToken(ctx context.Context, req *api.GetValueByAccessTokenRequest) (*api.GetValueResponseByAccessToken, error)
}

type TokenService interface {
	CheckAccessToken(accessToken string) (*helpers.Claims, error)
}

type Services struct {
	AuthService  AuthService
	UserService  UserService
	RoleService  RoleService
	TokenService TokenService
	DataService  DataService
}

type Dependencies struct {
	Repository *repository.Repositories
	Hasher     *helpers.Md5
	JWTManager *helpers.JWT
}

func NewServices(deps *Dependencies) *Services {
	tokeService := NewTokenService(deps.Hasher, deps.JWTManager)
	authService := NewAuthService(deps.Hasher, deps.JWTManager, deps.Repository.AuthRepository, deps.Repository.TokenRepository)
	userService := NewUserService(deps.Hasher, tokeService, deps.Repository.UserRepository, deps.Repository.AuthRepository, deps.Repository.RoleRepository)
	roleService := NewRoleService(deps.Hasher, tokeService, deps.Repository.RoleRepository, deps.Repository.AuthRepository)
	dataService := NewDataService(deps.Hasher, tokeService, deps.Repository.DataRepository, deps.Repository.RoleRepository, deps.Repository.AuthRepository)

	return &Services{
		AuthService:  authService,
		UserService:  userService,
		RoleService:  roleService,
		TokenService: tokeService,
		DataService:  dataService,
	}
}