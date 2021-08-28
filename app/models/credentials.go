package models

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Credentials struct {
	Login    string `json:"login"    db:"login"`
	Password string `json:"password" db:"password"`
}

func (c *Credentials) Validate() error {
	if len(c.Login) < 3 {
		return status.Error(codes.InvalidArgument, "login cannot be less than 3 characters long")
	}
	if len(c.Login) > 32 {
		return status.Error(codes.InvalidArgument, "length of login can not exceed 32 characters")
	}
	if len(c.Password) > 64 {
		return status.Error(codes.InvalidArgument, "length of email can not exceed 64 characters")
	}
	if len(c.Password) < 8 {
		return status.Error(codes.InvalidArgument, "password cannot be less than 8 characters long")
	}

	return nil
}
