package service

import (
	"context"
	"github.com/midaef/emmet-server/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	"github.com/midaef/emmet-server/internal/api"
	"github.com/midaef/emmet-server/internal/repository"
	"github.com/midaef/emmet-server/pkg/helpers"
)

type User struct {
	hasher         *helpers.Md5
	tokenManager   *helpers.JWT
	userRepository repository.UserRepository
	authRepository repository.AuthRepository
}

func NewUserService(hasher *helpers.Md5, tokenManager *helpers.JWT, userRepository repository.UserRepository,
	authRepository repository.AuthRepository) *User {
	return &User{
		hasher:         hasher,
		tokenManager:   tokenManager,
		userRepository: userRepository,
		authRepository: authRepository,
	}
}

func (s *User) CreateUserByAccessToken(ctx context.Context, req *api.CreateUserByAccessTokenRequest) (*api.CreateUserResponseByAccessToken, error) {
	claims, err := s.tokenManager.ParseJWT(req.AccessToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Token incorrect")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, status.Error(codes.Unauthenticated, "Token lifetime expired")
	}

	if claims.Subject != "root" {
		return nil, status.Error(codes.PermissionDenied, "Insufficient access rights")
	}

	if s.authRepository.IsExistByEmail(ctx, req.Login) {
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
