package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/midaef/emmet-server/internal/models"
)

type Token struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) *Token {
	return &Token{
		db: db,
	}
}

func (r *Token) Create(ctx context.Context, token *models.UserToken) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO tokens (token, exp, user_role, user_login) VALUES($1, $2, $3, $4)",
		token.AccessToken,
		token.Exp,
		token.Role,
		token.Login,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Token) IsExistAccessTokenByLogin(ctx context.Context, login string) bool {
	var id uint64
	r.db.QueryRowContext(ctx, "SELECT id FROM tokens WHERE user_login = $1", login).Scan(&id)
	if id == 0 {
		return false
	}

	return true
}

func (r *Token) DeleteByLogin(ctx context.Context, login string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM tokens WHERE user_login = $1", login)
	if err != nil {
		return err
	}

	return nil
}
