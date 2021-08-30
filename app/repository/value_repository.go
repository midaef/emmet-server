package repository

import "github.com/jmoiron/sqlx"

type Value struct {
	db *sqlx.DB
}

func NewValueRepository(db *sqlx.DB) *Value {
	return &Value{
		db: db,
	}
}
