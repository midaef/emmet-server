package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/midaef/emmet-server/app/models"
)

type Token struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) *Token {
	return &Token{
		db: db,
	}
}

func (t *Token) CreateToken(ctx context.Context, token *models.Token) error {
	_, err := t.db.ExecContext(ctx, "INSERT INTO tokens (user_id, access_token, refresh_token, exp) VALUES($1, $2, $3, $4)",
		token.UserID,
		token.AccessToken,
		token.RefreshToken,
		token.Exp,
	)
	if err != nil {
		return err
	}

	return nil
}

func (t *Token) GetTokenByAccessToken(ctx context.Context, accessToken string) (*models.Token, error) {
	var token models.Token
	err := t.db.GetContext(ctx, &token, "SELECT * FROM tokens WHERE access_token=$1", accessToken)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (t *Token) DeleteTokenByAccessToken(ctx context.Context, accessToken string) error {
	_, err := t.db.ExecContext(ctx, "DELETE FROM tokens WHERE access_token=$1", accessToken)
	if err != nil {
		return err
	}

	return nil
}
