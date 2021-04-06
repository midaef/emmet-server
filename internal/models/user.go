package models

type User struct {
	Login    string `db:"login"`
	Password string `db:"password"`
}