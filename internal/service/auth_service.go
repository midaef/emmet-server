package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/midaef/emmet-server/internal/api"
	"github.com/midaef/emmet-server/internal/models"
	"github.com/midaef/emmet-server/internal/repository"
	"github.com/midaef/emmet-server/pkg/helpers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth struct {
	hasher          *helpers.Md5
	tokenManager    *helpers.JWT
	authRepository  repository.AuthRepository
}

func NewAuthService(hasher *helpers.Md5, tokenManager *helpers.JWT, authRepository repository.AuthRepository) *Auth {
	return &Auth{
		hasher:          hasher,
		tokenManager:    tokenManager,
		authRepository:  authRepository,
	}
}

func (s *Auth) AuthWithCredentials(ctx context.Context, req *api.AuthWithCredentialsRequest) (*api.AuthResponseAccessToken, error) {
	if !s.authRepository.IsExistByEmail(ctx, req.Login) {
		return nil, status.Error(codes.NotFound, "Login not exists")
	}

	user := &models.User{
		Login: req.Login,
		Password: s.hasher.PasswordToMD5Hash(req.Password),
	}

	role, err := s.authRepository.GetUserRoleByLoginAndPassword(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Login or Password incorrect")
	}

	accessToken, err := s.tokenManager.CreateAccessToken(&helpers.Claims{
		Login: req.Login,
		StandardClaims: jwt.StandardClaims{
			Subject: role,
		},
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "Token creation error")
	}

	return &api.AuthResponseAccessToken{
		AccessToken: accessToken,
	}, nil
}