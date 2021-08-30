package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/midaef/emmet-server/app/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) IsExistByLogin(ctx context.Context, login string) bool {
	var id uint64

	err := u.db.GetContext(ctx, &id, "SELECT id FROM users WHERE login = $1", login)
	if err != nil {
		return false
	}

	if id == 0 {
		return false
	}

	return true
}

func (u *User) GetUserByCredentials(ctx context.Context, credentials *models.Credentials) (uint64, error) {
	var id uint64

	err := u.db.GetContext(ctx, &id, "SELECT id FROM users WHERE login = $1 AND password = $2", credentials.Login, credentials.Password)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, status.Errorf(codes.Internal, "user not found")
	}

	return id, nil
}

func (u *User) GetUserByUserID(ctx context.Context, userID uint64) (*models.User, error) {
	var user models.User

	err := u.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1", userID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) CreateUser(ctx context.Context, user *models.User) (uint64, error) {
	var id uint64
	err := u.db.QueryRowContext(ctx, "INSERT INTO users (login, password, user_role, created_by) VALUES ($1,$2,$3,$4) RETURNING id",
		user.Login,
		user.Password,
		user.Role,
		user.CreatedBy,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *User) GetUserIDByLogin(ctx context.Context, login string) (uint64, error) {
	var id uint64

	err := u.db.GetContext(ctx, &id, "SELECT id FROM users WHERE login = $1", login)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, status.Errorf(codes.Internal, "user not found")
	}

	return id, nil
}

func (u *User) DeleteUserByUserID(ctx context.Context, userID uint64) error {
	_, err := u.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) UpdateUserPasswordAndRoleByUserID(ctx context.Context, userID uint64, password string, role string) error {
	_, err := u.db.ExecContext(ctx, "UPDATE users SET (password, user_role) = ($1,$2) WHERE id = $3", password, role, userID)
	if err != nil {
		return err
	}

	return nil
}
