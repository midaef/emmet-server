package models

type User struct {
	Login    string `json:"login"    db:"login"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role"     db:"role"`
}
