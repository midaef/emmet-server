package models

type User struct {
	Login    string `db:"login"`
	Password string `db:"password"`
}

type CreateUser struct {
	AccessToken string `db:"token"`
	Login       string `db:"login"`
	Password    string `db:"password"`
	Role        string `db:"user_role"`
}