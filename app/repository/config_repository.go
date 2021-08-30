package repository

import "github.com/jmoiron/sqlx"

type Config struct {
	db *sqlx.DB
}

func NewConfigRepository(db *sqlx.DB) *Config {
	return &Config{
		db: db,
	}
}
