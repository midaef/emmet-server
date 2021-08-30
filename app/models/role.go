package models

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Role struct {
	ID           uint64  `json:"id"            db:"id"`
	RoleName     string  `json:"role_name"     db:"role_name"`
	CreatedBy    uint64  `json:"created_by"    db:"created_by"`
	CreateUser   bool    `json:"create_user"   db:"create_user"`
	DeleteUser   bool    `json:"delete_user"   db:"delete_user"`
	UpdateUser   bool    `json:"update_user"   db:"update_user"`
	CreateConfig bool    `json:"create_config" db:"create_config"`
	DeleteConfig bool    `json:"delete_config" db:"delete_config"`
	UpdateConfig bool    `json:"update_config" db:"update_config"`
	CreateRole   bool    `json:"create_role"   db:"create_role"`
	DeleteRole   bool    `json:"delete_role"   db:"delete_role"`
	UpdateRole   bool    `json:"update_role"   db:"update_role"`
	CreateValue  bool    `json:"create_value"  db:"create_value"`
	DeleteValue  bool    `json:"delete_value"  db:"delete_value"`
	UpdateValue  bool    `json:"update_value"  db:"update_value"`
	AllowedUsers []uint8 `json:"allowed_users" db:"allowed_users"`
}

func (r *Role) Validate() error {
	if len(r.RoleName) < 3 {
		return status.Error(codes.InvalidArgument, "role name cannot be less than 3 characters long")
	}
	if len(r.RoleName) > 32 {
		return status.Error(codes.InvalidArgument, "length of role name can not exceed 32 characters")
	}

	return nil
}
