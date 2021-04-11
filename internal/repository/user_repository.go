package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/midaef/emmet-server/internal/models"
)

type User struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *User {
	return &User{
		db: db,
	}
}

func (r *User) CreateUserByAccessToken(ctx context.Context, user *models.CreateUser) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (login, password, user_role) VALUES($1, $2, $3)",
		user.Login,
		user.Password,
		user.Role,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *User) DeleteUserByLogin(ctx context.Context, login string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE login = $1", login)
	if err != nil {
		return err
	}

	return nil
}

func (r *User) UpdatePasswordByLogin(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET password = $1 WHERE login = $2", user.Password, user.Login)
	if err != nil {
		return err
	}

	return nil
}

