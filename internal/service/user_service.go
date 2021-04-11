package service

import (
	"context"
	"github.com/midaef/emmet-server/internal/api"
	"github.com/midaef/emmet-server/internal/models"
	"github.com/midaef/emmet-server/internal/repository"
	"github.com/midaef/emmet-server/pkg/helpers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	hasher         *helpers.Md5
	tokenService   TokenService
	userRepository repository.UserRepository
	authRepository repository.AuthRepository
	roleRepository repository.RoleRepository
}

func NewUserService(hasher *helpers.Md5, tokenService TokenService, userRepository repository.UserRepository,
	authRepository repository.AuthRepository, roleRepository repository.RoleRepository) *User {
	return &User{
		hasher:         hasher,
		tokenService:   tokenService,
		userRepository: userRepository,
		authRepository: authRepository,
		roleRepository: roleRepository,
	}
}

func (s *User) CreateUserByAccessToken(ctx context.Context, req *api.CreateUserByAccessTokenRequest) (*api.CreateUserResponseByAccessToken, error) {
	claims, err := s.tokenService.CheckAccessToken(req.AccessToken)
	if err != nil {
		return nil, err
	}

	if !s.authRepository.IsExistByLogin(ctx, claims.Login) {
		return nil, status.Error(codes.NotFound, "Your account not exists")
	}

	permissions, err := s.roleRepository.GetPermissionsByRole(ctx, claims.Subject)
	if err != nil {
		return nil, status.Error(codes.Internal, "Get permissions error")
	}

	if !permissions.CreateUser {
		return nil, status.Error(codes.PermissionDenied, "Insufficient access rights")
	}

	if !s.roleRepository.IsExistByRole(ctx, req.Role) {
		return nil, status.Error(codes.NotFound, "Role not exists")
	}

	if s.authRepository.IsExistByLogin(ctx, req.Login) {
		return nil, status.Error(codes.AlreadyExists, "Login exists")
	}

	user := &models.CreateUser{
		AccessToken: req.AccessToken,
		Login:       req.Login,
		Password:    s.hasher.PasswordToMD5Hash(req.Password),
		Role:        req.Role,
	}

	err = s.userRepository.CreateUserByAccessToken(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, "User creation error")
	}

	return &api.CreateUserResponseByAccessToken{
		Message: "User created",
	}, nil
}

func (s *User) DeleteUserByAccessToken(ctx context.Context, req *api.DeleteUserByAccessTokenRequest) (*api.DeleteUserResponseByAccessToken, error) {
	claims, err := s.tokenService.CheckAccessToken(req.AccessToken)
	if err != nil {
		return nil, err
	}

	if !s.authRepository.IsExistByLogin(ctx, claims.Login) {
		return nil, status.Error(codes.NotFound, "Your account not exists")
	}

	permissions, err := s.roleRepository.GetPermissionsByRole(ctx, claims.Subject)
	if err != nil {
		return nil, status.Error(codes.Internal, "Get permissions error")
	}

	if !permissions.DeleteUser {
		return nil, status.Error(codes.PermissionDenied, "Insufficient access rights")
	}

	if req.Login == "root" {
		return nil, status.Error(codes.PermissionDenied, "Root cannot be deleted")
	}

	if !s.authRepository.IsExistByLogin(ctx, req.Login) {
		return nil, status.Error(codes.AlreadyExists, "Login not exists")
	}

	err = s.userRepository.DeleteUserByLogin(ctx, req.Login)
	if err != nil {
		return nil, status.Error(codes.Internal, "Delete user error")
	}

	return &api.DeleteUserResponseByAccessToken{
		Message: "User deleted",
	}, nil
}

func (s *User) UpdatePasswordByAccessToken(ctx context.Context, req *api.UpdatePasswordByAccessTokenRequest) (*api.UpdatePasswordResponseByAccessToken, error) {
	claims, err := s.tokenService.CheckAccessToken(req.AccessToken)
	if err != nil {
		return nil, err
	}

	if !s.authRepository.IsExistByLogin(ctx, claims.Login) {
		return nil, status.Error(codes.NotFound, "Your account not exists")
	}

	user := &models.User{
		Login:    claims.Login,
		Password: s.hasher.PasswordToMD5Hash(req.Password),
	}

	err = s.userRepository.UpdatePasswordByLogin(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Update password error")
	}

	return &api.UpdatePasswordResponseByAccessToken{
		Message: "Password updated",
	}, nil
}

