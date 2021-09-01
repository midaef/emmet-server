package user

import (
	"context"
	"github.com/midaef/emmet-server/extra/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *UserEndpoint) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.UserMessageResponse, error) {
	id, err := u.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), u.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	if !u.services.UserService.IsExistByLogin(ctx, req.GetUserLogin()) {
		return nil, status.Error(codes.NotFound, "login doesn't exist")
	}

	deleteUser, err := u.services.UserService.GetUserByLogin(ctx, req.GetUserLogin())
	if err != nil {
		return nil, err
	}

	if deleteUser.CreatedBy != id {
		userModel, err := u.services.UserService.GetUserByUserID(ctx, id)
		if err != nil {
			return nil, err
		}

		roleModel, err := u.services.RoleService.GetRoleIDByName(ctx, userModel.Role)
		if err != nil {
			return nil, err
		}

		if !roleModel.DeleteUser {
			return nil, status.Error(codes.PermissionDenied, "your role does not have enough rights to delete a user")
		}
	}

	err = u.services.UserService.DeleteUserByUserID(ctx, deleteUser.ID)
	if err != nil {
		return nil, err
	}

	return &user.UserMessageResponse{
		Message: "user was deleted successfully",
	}, nil
}
