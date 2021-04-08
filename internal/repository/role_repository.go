package repository

import "github.com/jmoiron/sqlx"

type Role struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) *Role {
	return &Role{
		db: db,
	}
}
