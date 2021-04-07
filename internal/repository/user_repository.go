package repository

import "github.com/jmoiron/sqlx"

type User struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *User {
	return &User{
		db: db,
	}
}
