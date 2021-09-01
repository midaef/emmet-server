package service

import (
	"context"
	"github.com/midaef/emmet-server/app/models"
	"github.com/midaef/emmet-server/tools/helpers"
)

func (s *Service) IsExistByRole(ctx context.Context, role string) bool {
	return s.repository.RoleRepository.IsExistByRole(ctx, role)
}

func (s *Service) GetRoleIDByName(ctx context.Context, name string) (*models.Role, error) {
	id, err := s.repository.RoleRepository.GetRoleIDByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return s.GetRoleByRoleID(ctx, id)
}

func (s *Service) GetRoleByRoleID(ctx context.Context, roleID uint64) (*models.Role, error) {
	return s.repository.RoleRepository.GetRoleByRoleID(ctx, roleID)
}

func (s *Service) IsRoleAllowedForUser(ctx context.Context, userID uint64, role string) (bool, error) {
	roleModelCreationUser, err := s.GetRoleIDByName(ctx, role)
	if err != nil {
		return false, err
	}

	isAllowed := false

	allowedUsers, err := helpers.StringToUint64Array(string(roleModelCreationUser.AllowedUsers))
	if err != nil {
		return false, err
	}

	if len(allowedUsers) == 0 {
		return false, nil
	}

	for _, idLocal := range allowedUsers {
		if idLocal == userID {
			isAllowed = true
		}
	}

	return isAllowed, nil
}

func (s *Service) CreateRole(ctx context.Context, role *models.Role) (uint64, error) {
	return s.repository.RoleRepository.CreateRole(ctx, role)
}
