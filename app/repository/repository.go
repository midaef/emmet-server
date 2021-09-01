package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/midaef/emmet-server/app/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (uint64, error)
	IsExistByLogin(ctx context.Context, login string) bool
	GetUserByCredentials(ctx context.Context, credentials *models.Credentials) (uint64, error)
	GetUserByUserID(ctx context.Context, userID uint64) (*models.User, error)
	GetUserIDByLogin(ctx context.Context, login string) (uint64, error)
	DeleteUserByUserID(ctx context.Context, userID uint64) error
	UpdateUserPasswordAndRoleByUserID(ctx context.Context, userID uint64, password string, role string) error
}

type TokenRepository interface {
	CreateToken(ctx context.Context, token *models.Token) error
	GetTokenByAccessToken(ctx context.Context, accessToken string) (*models.Token, error)
	DeleteTokenByAccessToken(ctx context.Context, accessToken string) error
}

type RoleRepository interface {
	IsExistByRole(ctx context.Context, role string) bool
	GetRoleIDByName(ctx context.Context, name string) (uint64, error)
	GetRoleByRoleID(ctx context.Context, roleID uint64) (*models.Role, error)
	CreateRole(ctx context.Context, role *models.Role) (uint64, error)
}

type ConfigRepository interface {
}

type ValueRepository interface {
}

type Repository struct {
	UserRepository   UserRepository
	TokenRepository  TokenRepository
	RoleRepository   RoleRepository
	ConfigRepository ConfigRepository
	ValueRepository  ValueRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:   NewUserRepository(db),
		TokenRepository:  NewTokenRepository(db),
		RoleRepository:   NewRoleRepository(db),
		ConfigRepository: NewConfigRepository(db),
		ValueRepository:  NewValueRepository(db),
	}
}
