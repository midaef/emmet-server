package models

import (
	"time"
)

type Token struct {
	AccessToken  string    `json:"access_token"  db:"access_token"`
	RefreshToken string    `json:"refresh_token" db:"refresh"`
	Exp          time.Time `json:"exp"           db:"exp"`
	UserID       uint64    `json:"user_id"       db:"user_id"`
}
