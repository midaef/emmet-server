package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/midaef/emmet-server/internal/models"
)

type Role struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) *Role {
	return &Role{
		db: db,
	}
}

func (r *Role) CreateRole(ctx context.Context, role *models.Role) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO roles (created_by, create_user, create_role, create_value, user_role) " +
		"VALUES($1, $2, $3, $4, $5)",
		role.CreatedBy,
		role.CreateUser,
		role.CreateRole,
		role.CreateValue,
		role.Role,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Role) GetPermissionsByRole(ctx context.Context, role string) (*models.Permissions, error) {
	var permissions *models.Permissions
	err := r.db.GetContext(ctx, &permissions, "SELECT create_user, create_role, create_value FROM roles WHERE user_role = $1", role)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}