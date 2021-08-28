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
	var id []uint64

	err := u.db.GetContext(ctx, &id, "SELECT id FROM users WHERE login = $1 AND password = $2", credentials.Login, credentials.Password)
	if err != nil {
		return 0, err
	}

	if len(id) == 0 {
		return 0, status.Errorf(codes.Internal, "user not found")
	}

	return id[0], nil
}

func (u *User) GetUserByUserID(ctx context.Context, userID uint64) (string, error) {
	var userType []string

	err := u.db.SelectContext(ctx, &userType, "SELECT user_type FROM users WHERE id=$1", userID)
	if err != nil {
		return "", err
	}

	if len(userType) == 0 {
		return "", status.Errorf(codes.Internal, "User not found")
	}

	return userType[0], err
}

func (u *User) CreateUser(ctx context.Context, user *models.User) (uint64, error) {
	var id uint64
	err := u.db.QueryRowContext(ctx, "INSERT INTO users (login, password, role) VALUES ($1,$2,$3) RETURNING id",
		user.Login,
		user.Password,
		user.Role,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
