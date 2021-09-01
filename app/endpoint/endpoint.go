package endpoint

import (
	"context"
	"github.com/midaef/emmet-server/extra/auth"
	"github.com/midaef/emmet-server/extra/role"
	"github.com/midaef/emmet-server/extra/user"
)

type EndpointContainer struct {
	AuthService AuthServiceInter
	UserService UserServiceInter
	RoleService RoleServiceInter
}

func NewEndpointContainer(auth AuthServiceInter, user UserServiceInter, role RoleServiceInter) *EndpointContainer {
	return &EndpointContainer{
		AuthService: auth,
		UserService: user,
		RoleService: role,
	}
}

type AuthServiceInter interface {
	AuthWithCredentials(ctx context.Context, req *auth.AuthWithCredentialsRequest) (*auth.AccessRefreshTokens, error)
	RefreshTokens(ctx context.Context, req *auth.AccessRefreshTokens) (*auth.AccessRefreshTokens, error)
	AuthWithAccessToken(ctx context.Context, req *auth.AuthWithAccessTokenRequest) (*auth.AuthWithAccessTokenResponse, error)
}

type UserServiceInter interface {
	CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserMessageResponse, error)
	UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UserMessageResponse, error)
	DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.UserMessageResponse, error)
}

type RoleServiceInter interface {
	CreateRole(ctx context.Context, req *role.RoleRequest) (*role.RoleMessageResponse, error)
	UpdateRole(ctx context.Context, req *role.RoleRequest) (*role.RoleMessageResponse, error)
	DeleteRole(ctx context.Context, req *role.DeleteRoleRequest) (*role.RoleMessageResponse, error)
}
