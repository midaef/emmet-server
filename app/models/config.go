package models

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Config struct {
	ID           uint64   `json:"id"            db:"id"`
	ConfigName   string   `json:"config_name"   db:"config_name"`
	CreatedBy    uint64   `json:"created_by"    db:"created_by"`
	AllowedRoles []uint64 `json:"allowed_roles" db:"allowed_roles"`
}

func (c *Config) Validate() error {
	if len(c.ConfigName) < 3 {
		return status.Error(codes.InvalidArgument, "config name cannot be less than 3 characters long")
	}
	if len(c.ConfigName) > 32 {
		return status.Error(codes.InvalidArgument, "length of config name can not exceed 32 characters")
	}

	return nil
}
