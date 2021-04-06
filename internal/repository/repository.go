package repository

import "github.com/jmoiron/sqlx"

type Repositories struct {

}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{}
}