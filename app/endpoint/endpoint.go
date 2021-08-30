package endpoint

import (
	"context"
	"github.com/midaef/emmet-server/extra/auth"
	"github.com/midaef/emmet-server/extra/user"
)

type EndpointContainer struct {
	AuthService AuthServiceInter
	UserService UserServiceInter
}

func NewEndpointContainer(auth AuthServiceInter, user UserServiceInter) *EndpointContainer {
	return &EndpointContainer{
		AuthService: auth,
		UserService: user,
	}
}

type AuthServiceInter interface {
	AuthWithCredentials(ctx context.Context, req *auth.AuthWithCredentialsRequest) (*auth.AccessRefreshTokens, error)
	RefreshTokens(ctx context.Context, req *auth.AccessRefreshTokens) (*auth.AccessRefreshTokens, error)
	AuthWithAccessToken(ctx context.Context, req *auth.AuthWithAccessTokenRequest) (*auth.AuthWithAccessTokenResponse, error)
}

type UserServiceInter interface {
	CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.MessageResponse, error)
	UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.MessageResponse, error)
	DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.MessageResponse, error)
}
