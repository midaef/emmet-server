package user

import (
	"context"
	"github.com/midaef/emmet-server/app/models"
	"github.com/midaef/emmet-server/extra/user"
	"github.com/midaef/emmet-server/tools/helpers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *UserEndpoint) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.MessageResponse, error) {
	id, err := u.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), u.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	if u.services.UserService.IsExistByLogin(ctx, req.GetUserLogin()) {
		return nil, status.Error(codes.NotFound, "login is exist")
	}

	if !u.services.RoleService.IsExistByRole(ctx, req.GetUserRole()) {
		return nil, status.Error(codes.NotFound, "role doesn't exist")
	}

	userModel, err := u.services.UserService.GetUserByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	roleModel, err := u.services.RoleService.GetRoleIDByName(ctx, userModel.Role)
	if err != nil {
		return nil, err
	}

	if !roleModel.CreateUser {
		return nil, status.Error(codes.PermissionDenied, "your role does not have enough rights to create a user")
	}

	isAllowed, err := u.services.RoleService.IsRoleAllowedForUser(ctx, id, req.GetUserRole())
	if err != nil {
		return nil, err
	}

	if !isAllowed {
		return nil, status.Error(codes.PermissionDenied, "there are not enough rights to assign this role to the user")
	}

	newUser := &models.User{
		Login:     req.GetUserLogin(),
		Password:  req.GetUserPassword(),
		Role:      req.GetUserRole(),
		CreatedBy: id,
	}

	if err = newUser.Validate(); err != nil {
		return nil, err
	}

	hashPassword, err := helpers.NewMD5Hash(req.GetUserPassword())
	if err != nil {
		return nil, err
	}

	newUser.Password = hashPassword

	_, err = u.services.UserService.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &user.MessageResponse{
		Message: "user was created successfully",
	}, nil
}
