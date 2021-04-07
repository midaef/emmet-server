package models

import "time"

type UserToken struct {
	AccessToken  string    `json:"access_token"`
	Exp          time.Time `json:"exp"`
	Role         string    `json:"role"`
	Login        string    `json:"login"`
}