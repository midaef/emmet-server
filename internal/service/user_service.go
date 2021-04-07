package service

import (
	"github.com/midaef/emmet-server/internal/repository"
	"github.com/midaef/emmet-server/pkg/helpers"
)

type User struct {
	hasher          *helpers.Md5
	tokenManager    *helpers.JWT
	userRepository  repository.UserRepository
}

func NewUserService(hasher *helpers.Md5, tokenManager *helpers.JWT, userRepository repository.UserRepository) *User {
	return &User{
		hasher:          hasher,
		tokenManager:    tokenManager,
		userRepository: userRepository,
	}
}
