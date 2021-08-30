package user

import (
	"context"
	"github.com/midaef/emmet-server/app/models"
	"github.com/midaef/emmet-server/extra/user"
	"github.com/midaef/emmet-server/tools/helpers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *UserEndpoint) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.MessageResponse, error) {
	id, err := u.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), u.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	if !u.services.UserService.IsExistByLogin(ctx, req.GetUserLogin()) {
		return nil, status.Error(codes.NotFound, "login doesn't' exist")
	}

	updateUser, err := u.services.UserService.GetUserByLogin(ctx, req.GetUserLogin())
	if err != nil {
		return nil, err
	}

	if updateUser.CreatedBy != id {
		userModel, err := u.services.UserService.GetUserByUserID(ctx, id)
		if err != nil {
			return nil, err
		}

		roleModel, err := u.services.RoleService.GetRoleIDByName(ctx, userModel.Role)
		if err != nil {
			return nil, err
		}

		if !roleModel.UpdateUser {
			return nil, status.Error(codes.PermissionDenied, "your role does not have enough rights to update a user")
		}
	}

	var role, password string

	if req.GetNewUserPassword() == "" {
		password = updateUser.Password
	} else {
		password = req.GetNewUserPassword()

		if err = models.ValidatePassword(req.GetNewUserPassword()); err != nil {
			return nil, err
		}

		hashPassword, err := helpers.NewMD5Hash(req.GetNewUserPassword())
		if err != nil {
			return nil, err
		}

		password = hashPassword
	}

	if req.GetNewUserRole() == "" {
		role = updateUser.Role
	} else {
		role = req.GetNewUserRole()

		if err = models.ValidateRole(req.GetNewUserRole()); err != nil {
			return nil, err
		}

		if !u.services.RoleService.IsExistByRole(ctx, req.GetNewUserRole()) {
			return nil, status.Error(codes.NotFound, "role doesn't exist")
		}

		isAllowed, err := u.services.RoleService.IsRoleAllowedForUser(ctx, id, req.GetNewUserRole())
		if err != nil {
			return nil, err
		}

		if !isAllowed {
			return nil, status.Error(codes.PermissionDenied, "there are not enough rights to assign this role to the user")
		}
	}

	err = u.services.UserService.UpdateUserPasswordAndRoleByUserID(ctx, updateUser.ID, password, role)
	if err != nil {
		return nil, err
	}

	return &user.MessageResponse{
		Message: "user was updated successfully",
	}, nil
}
