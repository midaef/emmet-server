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

type Role struct {
	hasher         *helpers.Md5
	tokenService   TokenService
	roleRepository repository.RoleRepository
}

func NewRoleService(hasher *helpers.Md5, tokenService TokenService, roleRepository repository.RoleRepository) *Role {
	return &Role{
		hasher:         hasher,
		tokenService:   tokenService,
		roleRepository: roleRepository,
	}
}

func (s *Role) CreateRoleByAccessToken(ctx context.Context, req *api.CreateRoleByAccessTokenRequest) (*api.CreateRoleResponseByAccessToken, error) {
	claims, err := s.tokenService.CheckAccessToken(req.AccessToken)
	if err != nil {
		return nil, err
	}

	permissions, err := s.roleRepository.GetPermissionsByRole(ctx, claims.Subject)
	if err != nil {
		return nil, status.Error(codes.Internal, "Get permissions error")
	}

	if !permissions.CreateRole {
		return nil, status.Error(codes.PermissionDenied, "Insufficient access rights")
	}

	if s.roleRepository.IsExistByRole(ctx, req.Role) {
		return nil, status.Error(codes.AlreadyExists, "Role exists")
	}

	role := &models.Role{
		CreatedBy:   claims.Subject,
		CreateRole:  req.CreateRole,
		CreateUser:  req.CreateUser,
		CreateValue: req.CreateValue,
		DeleteRole:  req.DeleteRole,
		DeleteUser:  req.DeleteUser,
		DeleteValue: req.DeleteValue,
		Role:        req.Role,
	}

	err = s.roleRepository.CreateRole(ctx, role)
	if err != nil {
		return nil, status.Error(codes.Internal, "Create role error")
	}

	return &api.CreateRoleResponseByAccessToken{
		Message: "Role created",
	}, nil
}

func (s *Role) DeleteRoleByAccessToken(ctx context.Context, req *api.DeleteRoleByAccessTokenRequest) (*api.DeleteRoleResponseByAccessToken, error) {
	claims, err := s.tokenService.CheckAccessToken(req.AccessToken)
	if err != nil {
		return nil, err
	}

	permissions, err := s.roleRepository.GetPermissionsByRole(ctx, claims.Subject)
	if err != nil {
		return nil, status.Error(codes.Internal, "Get permissions error")
	}

	if !permissions.DeleteRole {
		return nil, status.Error(codes.PermissionDenied, "Insufficient access rights")
	}

	if !s.roleRepository.IsExistByRole(ctx, req.Role) {
		return nil, status.Error(codes.AlreadyExists, "Role not exists")
	}

	err = s.roleRepository.DeleteByRole(ctx, req.Role)
	if err != nil {
		return nil, status.Error(codes.Internal, "Delete role error")
	}

	return &api.DeleteRoleResponseByAccessToken{
		Message: "Role deleted",
	}, nil
}