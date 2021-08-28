package service

import (
	"context"
	"github.com/midaef/emmet-server/app/models"
	"github.com/midaef/emmet-server/app/repository"
	"github.com/midaef/emmet-server/config"
	"time"
)

type ITokenService interface {
	CreateToken(ctx context.Context, token *models.Token) error
	GenerateTokens(ctx context.Context, userID uint64, jwtSecretKey string, expirationTime time.Duration) (*models.Token, error)
	GetTokenByAccessToken(ctx context.Context, accessToken string) (*models.Token, error)
	DeleteTokenByAccessToken(ctx context.Context, accessToken string) error
}

type IUserService interface {
	CreateUser(ctx context.Context, user *models.User) (uint64, error)

	IsExistByLogin(ctx context.Context, login string) bool

	GetUserByCredentials(ctx context.Context, credentials *models.Credentials) (uint64, error)
	GetUserByUserID(ctx context.Context, userID uint64) (string, error)
}

type Services struct {
	TokenService ITokenService
	UserService  IUserService
}

type Service struct {
	repository *repository.Repository
	config     *config.Config
	Services   *Services
}

func NewService(repository *repository.Repository, config *config.Config) *Service {
	return &Service{
		repository: repository,
		config:     config,
	}
}
