package models

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	ID        uint64 `json:"id"         db:"id"`
	Login     string `json:"login"      db:"login"`
	Password  string `json:"password"   db:"password"`
	Role      string `json:"user_role"  db:"user_role"`
	CreatedBy uint64 `json:"created_by" db:"created_by"`
}

func (u *User) Validate() error {
	if len(u.Login) < 3 {
		return status.Error(codes.InvalidArgument, "login cannot be less than 3 characters long")
	}
	if len(u.Login) > 32 {
		return status.Error(codes.InvalidArgument, "length of login can not exceed 32 characters")
	}
	if len(u.Password) > 64 {
		return status.Error(codes.InvalidArgument, "length of email can not exceed 64 characters")
	}
	if len(u.Password) < 8 {
		return status.Error(codes.InvalidArgument, "password cannot be less than 8 characters long")
	}
	if len(u.Role) < 3 {
		return status.Error(codes.InvalidArgument, "role name cannot be less than 3 characters long")
	}
	if len(u.Role) > 32 {
		return status.Error(codes.InvalidArgument, "length of role name can not exceed 32 characters")
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) > 64 {
		return status.Error(codes.InvalidArgument, "length of email can not exceed 64 characters")
	}
	if len(password) < 8 {
		return status.Error(codes.InvalidArgument, "password cannot be less than 8 characters long")
	}

	return nil
}

func ValidateRole(role string) error {
	if len(role) < 3 {
		return status.Error(codes.InvalidArgument, "role name cannot be less than 3 characters long")
	}
	if len(role) > 32 {
		return status.Error(codes.InvalidArgument, "length of role name can not exceed 32 characters")
	}

	return nil
}
