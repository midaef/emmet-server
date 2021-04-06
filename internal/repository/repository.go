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

type Repositories struct {
	AuthRepository AuthRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		AuthRepository: NewAuthRepository(db),
	}
}