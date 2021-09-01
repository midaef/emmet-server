package role

import (
	"context"
	"github.com/midaef/emmet-server/app/models"
	"github.com/midaef/emmet-server/extra/role"
	"github.com/midaef/emmet-server/tools/helpers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *RoleEndpoint) CreateRole(ctx context.Context, req *role.RoleRequest) (*role.RoleMessageResponse, error) {
	id, err := r.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), r.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	if r.services.RoleService.IsExistByRole(ctx, req.GetRoleName()) {
		return nil, status.Error(codes.NotFound, "role is exist")
	}

	userModel, err := r.services.UserService.GetUserByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	roleModel, err := r.services.RoleService.GetRoleIDByName(ctx, userModel.Role)
	if err != nil {
		return nil, err
	}

	if !roleModel.CreateRole {
		return nil, status.Error(codes.PermissionDenied, "your role does not have enough rights to create a user")
	}

	allowedID := make([]uint64, 0, len(req.GetAllowedUsers()))

	for _, login := range req.GetAllowedUsers() {
		if !r.services.UserService.IsExistByLogin(ctx, login) {
			return nil, status.Error(codes.NotFound, login+" (user) doesn't exist")
		}

		idLocal, err := r.services.UserService.GetUserIDByLogin(ctx, login)
		if err != nil {
			return nil, err
		}

		allowedID = append(allowedID, idLocal)
	}

	allowedUsers, err := helpers.Uint64ArrayToString(allowedID)
	if err != nil {
		return nil, err
	}

	roleModelCreation := &models.Role{
		RoleName:     req.GetRoleName(),
		CreatedBy:    id,
		CreateUser:   req.GetCreateUser(),
		DeleteUser:   req.GetDeleteUser(),
		UpdateUser:   req.GetUpdateUser(),
		CreateConfig: req.GetCreateConfig(),
		DeleteConfig: req.GetDeleteConfig(),
		UpdateConfig: req.GetUpdateConfig(),
		CreateRole:   req.GetCreateRole(),
		DeleteRole:   req.GetDeleteRole(),
		UpdateRole:   req.GetUpdateRole(),
		CreateValue:  req.GetCreateValue(),
		DeleteValue:  req.GetDeleteValue(),
		UpdateValue:  req.GetUpdateValue(),
		AllowedUsers: []byte(allowedUsers),
	}

	if err = roleModelCreation.Validate(); err != nil {
		return nil, err
	}

	_, err = r.services.RoleService.CreateRole(ctx, roleModelCreation)
	if err != nil {
		return nil, err
	}

	return &role.RoleMessageResponse{
		Message: "role was created successfully",
	}, nil
}
