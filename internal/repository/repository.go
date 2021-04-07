package repository

import (
	"context"
	"github.com/midaef/emmet-server/internal/models"

	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	IsExistByEmail(ctx context.Context, login string) bool
	GetUserRoleByLoginAndPassword(ctx context.Context, user *models.User) (string, error)
}

type TokenRepository interface {
	Create(ctx context.Context, token *models.UserToken) error
	IsExistAccessTokenByLogin(ctx context.Context, login string) bool
	DeleteByLogin(ctx context.Context, login string) error
}

type UserRepository interface {

}

type Repositories struct {
	AuthRepository  AuthRepository
	TokenRepository TokenRepository
	UserRepository  UserRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		AuthRepository: NewAuthRepository(db),
		TokenRepository: NewTokenRepository(db),
		UserRepository: NewUserRepository(db),
	}
}