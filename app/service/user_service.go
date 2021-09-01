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

func (s *Service) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	id, err := s.repository.UserRepository.GetUserIDByLogin(ctx, login)
	if err != nil {
		return nil, err
	}

	return s.GetUserByUserID(ctx, id)
}

func (s *Service) GetUserByUserID(ctx context.Context, userID uint64) (*models.User, error) {
	return s.repository.UserRepository.GetUserByUserID(ctx, userID)
}

func (s *Service) CreateUser(ctx context.Context, user *models.User) (uint64, error) {
	return s.repository.UserRepository.CreateUser(ctx, user)
}

func (s *Service) DeleteUserByUserID(ctx context.Context, userID uint64) error {
	return s.repository.UserRepository.DeleteUserByUserID(ctx, userID)
}

func (s *Service) UpdateUserPasswordAndRoleByUserID(ctx context.Context, userID uint64, password string, role string) error {
	return s.repository.UserRepository.UpdateUserPasswordAndRoleByUserID(ctx, userID, password, role)
}

func (s *Service) GetUserIDByLogin(ctx context.Context, login string) (uint64, error) {
	return s.repository.UserRepository.GetUserIDByLogin(ctx, login)
}
