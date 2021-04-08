package service

import (
	"context"
	"github.com/midaef/emmet-server/internal/api"

	"github.com/midaef/emmet-server/pkg/helpers"
)

type Role struct {
	hasher         *helpers.Md5
	tokenManager   *helpers.JWT
}

func NewRoleService(hasher *helpers.Md5, tokenManager *helpers.JWT) *Role {
	return &Role{
		hasher:         hasher,
		tokenManager:   tokenManager,
	}
}

func (s *Role) CreateRoleByAccessToken(ctx context.Context, req *api.CreateRoleByAccessTokenRequest) (*api.CreateRoleResponseByAccessToken, error) {
	return &api.CreateRoleResponseByAccessToken{
		Message: "Role created",
	}, nil
}