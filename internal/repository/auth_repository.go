package repository

import (
	"context"
	"github.com/midaef/emmet-server/internal/models"

	"github.com/jmoiron/sqlx"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *Auth {
	return &Auth{
		db: db,
	}
}

func (r *Auth) IsExistByLogin(ctx context.Context, login string) bool {
	var id uint64
	r.db.QueryRowContext(ctx, "SELECT id FROM users WHERE login = $1", login).Scan(&id)
	if id == 0 {
		return false
	}

	return true
}

func (r *Auth) GetUserRoleByLoginAndPassword(ctx context.Context, user *models.User) (string, error) {
	var role string
	err := r.db.GetContext(ctx, &role, "SELECT user_role FROM users WHERE login=$1 AND password=$2", user.Login, user.Password)
	if err != nil {
		return "", err
	}

	return role, nil
}