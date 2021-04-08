package repository

import (
	"context"
	"github.com/midaef/emmet-server/internal/models"

	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	IsExistByLogin(ctx context.Context, login string) bool
	GetUserRoleByLoginAndPassword(ctx context.Context, user *models.User) (string, error)
}

type TokenRepository interface {
	Create(ctx context.Context, token *models.UserToken) error
	IsExistAccessTokenByLogin(ctx context.Context, login string) bool
	DeleteByLogin(ctx context.Context, login string) error
}

type UserRepository interface {
	CreateUserByAccessToken(ctx context.Context, user *models.CreateUser) error
	DeleteUserByLogin(ctx context.Context, login string) error
}

type RoleRepository interface {
	CreateRole(ctx context.Context, role *models.Role) error
	GetPermissionsByRole(ctx context.Context, role string) (*models.Permissions, error)
	IsExistByRole(ctx context.Context, role string) bool
	DeleteByRole(ctx context.Context, role string) error
}

type Repositories struct {
	AuthRepository  AuthRepository
	TokenRepository TokenRepository
	UserRepository  UserRepository
	RoleRepository  RoleRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		AuthRepository: NewAuthRepository(db),
		TokenRepository: NewTokenRepository(db),
		UserRepository: NewUserRepository(db),
		RoleRepository: NewRoleRepository(db),
	}
}