package service

import (
	"context"
	"github.com/midaef/emmet-server/app/models"
)

func (s *Service) IsExistByLogin(ctx context.Context, login string) bool {
	return s.repository.UserRepository.IsExistByLogin(ctx, login)
}

func (s *Service) GetUserByCredentials(ctx context.Context, credentials *models.Credentials) (uint64, error) {
	return s.repository.UserRepository.GetUserByCredentials(ctx, credentials)
}

func (s *Service) GetUserByUserID(ctx context.Context, userID uint64) (string, error) {
	return s.repository.UserRepository.GetUserByUserID(ctx, userID)
}

func (s *Service) CreateUser(ctx context.Context, user *models.User) (uint64, error) {
	return s.repository.UserRepository.CreateUser(ctx, user)
}
