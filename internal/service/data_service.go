package service

import (
	"context"
	"github.com/midaef/emmet-server/internal/models"
	"strings"

	"github.com/midaef/emmet-server/internal/api"
	"github.com/midaef/emmet-server/internal/repository"
	"github.com/midaef/emmet-server/pkg/helpers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Data struct {
	hasher         *helpers.Md5
	tokenService   TokenService
	dataRepository repository.DataRepository
	roleRepository repository.RoleRepository
	authRepository repository.AuthRepository
}

func NewDataService(hasher *helpers.Md5, tokenService TokenService, dataRepository repository.DataRepository,
	roleRepository repository.RoleRepository, authRepository repository.AuthRepository) *Data {
	return &Data{
		hasher:         hasher,
		tokenService:   tokenService,
		dataRepository: dataRepository,
		roleRepository: roleRepository,
		authRepository: authRepository,
	}
}

func (s *Data) CreateValueByAccessToken(ctx context.Context, req *api.CreateValueByAccessTokenRequest) (*api.CreateValueResponseByAccessToken, error) {
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

	if !permissions.CreateValue {
		return nil, status.Error(codes.PermissionDenied, "Insufficient access rights")
	}

	if s.dataRepository.IsExistByKey(ctx, req.Key) {
		return nil, status.Error(codes.AlreadyExists, "Key exists")
	}

	value := &models.Value{
		CreatedBy: claims.Login,
		Key:       req.Key,
		Value:     req.Value,
		Roles:     req.Roles,
	}

	err = s.dataRepository.CreateValue(ctx, value)
	if err != nil {
		return nil, status.Error(codes.Internal, "Create value error")
	}

	return &api.CreateValueResponseByAccessToken{
		Message: "Value created",
	}, nil
}

func (s *Data) DeleteValueByAccessToken(ctx context.Context, req *api.DeleteValueByAccessTokenRequest) (*api.DeleteValueResponseByAccessToken, error) {
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

	if !permissions.DeleteValue {
		return nil, status.Error(codes.PermissionDenied, "Insufficient access rights")
	}

	if !s.dataRepository.IsExistByKey(ctx, req.Key) {
		return nil, status.Error(codes.AlreadyExists, "Key not exists")
	}

	err = s.dataRepository.DeleteValueByKey(ctx, req.Key)
	if err != nil {
		return nil, status.Error(codes.Internal, "Delete value error")
	}

	return &api.DeleteValueResponseByAccessToken{
		Message: "Value deleted",
	}, nil
}

func (s *Data) GetValueByAccessToken(ctx context.Context, req *api.GetValueByAccessTokenRequest) (*api.GetValueResponseByAccessToken, error) {
	claims, err := s.tokenService.CheckAccessToken(req.AccessToken)
	if err != nil {
		return nil, err
	}

	if !s.authRepository.IsExistByLogin(ctx, claims.Login) {
		return nil, status.Error(codes.NotFound, "Your account not exists")
	}

	if !s.dataRepository.IsExistByKey(ctx, req.Key) {
		return nil, status.Error(codes.AlreadyExists, "Key not exists")
	}

	value, err := s.dataRepository.GetValueByKey(ctx, req.Key)
	if err != nil {
		return nil, status.Error(codes.Internal, "Get value error")
	}

	roles := strings.Split(value.Roles, ",")
	for _, r := range roles {
		if r == claims.Subject || claims.Subject == "root" || value.CreatedBy == claims.Login {
			return &api.GetValueResponseByAccessToken{
				Value: value.Value,
			}, nil
		}
	}

	return nil, status.Error(codes.PermissionDenied, "Insufficient access rights")
}

