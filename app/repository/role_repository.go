package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/midaef/emmet-server/app/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Role struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) *Role {
	return &Role{
		db: db,
	}
}

func (r *Role) IsExistByRole(ctx context.Context, role string) bool {
	var id uint64

	err := r.db.GetContext(ctx, &id, "SELECT id FROM roles WHERE role_name = $1", role)
	if err != nil {
		return false
	}

	if id == 0 {
		return false
	}

	return true
}

func (r *Role) GetRoleIDByName(ctx context.Context, name string) (uint64, error) {
	var id uint64

	err := r.db.GetContext(ctx, &id, "SELECT id FROM roles WHERE role_name = $1", name)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, status.Error(codes.NotFound, "role doesn't exist")
	}

	return id, nil
}

func (r *Role) GetRoleByRoleID(ctx context.Context, roleID uint64) (*models.Role, error) {
	var role models.Role

	err := r.db.GetContext(ctx, &role, "SELECT * FROM roles WHERE id = $1", roleID)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *Role) CreateRole(ctx context.Context, role *models.Role) (uint64, error) {
	var id uint64

	err := r.db.QueryRowContext(ctx, "INSERT INTO roles (role_name, create_user, delete_user, update_user, "+
		"create_config, delete_config, update_config, create_role, delete_role, update_role, create_value, delete_value, "+
		"update_value, created_by, allowed_users) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) RETURNING id",
		role.RoleName,
		role.CreateUser,
		role.DeleteUser,
		role.UpdateUser,
		role.CreateConfig,
		role.DeleteConfig,
		role.UpdateConfig,
		role.CreateRole,
		role.DeleteRole,
		role.UpdateRole,
		role.CreateValue,
		role.DeleteValue,
		role.UpdateValue,
		role.CreatedBy,
		role.AllowedUsers,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
