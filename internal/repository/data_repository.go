package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/midaef/emmet-server/internal/models"
)

type Data struct {
	db *sqlx.DB
}

func NewDataRepository(db *sqlx.DB) *Data {
	return &Data{
		db: db,
	}
}

func (r *Data) CreateValue(ctx context.Context, value *models.Value) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO values (created_by, key, value, roles) VALUES($1, $2, $3, $4)",
		value.CreatedBy,
		value.Key,
		value.Value,
		value.Roles,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Data) DeleteValueByKey(ctx context.Context, key string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM values WHERE key = $1", key)
	if err != nil {
		return err
	}

	return nil
}

func (r *Data) GetValueByKey(ctx context.Context, key string) (*models.Value, error) {
	var value models.Value
	err := r.db.GetContext(ctx, &value, "SELECT created_by, key, value, roles FROM values WHERE key = $1", key)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (r *Data) IsExistByKey(ctx context.Context, key string) bool {
	var id uint64
	r.db.QueryRowContext(ctx, "SELECT id FROM values WHERE key = $1", key).Scan(&id)
	if id == 0 {
		return false
	}

	return true
}