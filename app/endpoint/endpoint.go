package endpoint

import (
	"context"
	"github.com/midaef/emmet-server/extra/auth"
)

type EndpointContainer struct {
	AuthService AuthServiceInter
}

func NewEndpointContainer(auth AuthServiceInter) *EndpointContainer {
	return &EndpointContainer{
		AuthService: auth,
	}
}

type AuthServiceInter interface {
	AuthWithCredentials(ctx context.Context, req *auth.AuthWithCredentialsRequest) (*auth.AccessRefreshTokens, error)

	RefreshTokens(ctx context.Context, req *auth.AccessRefreshTokens) (*auth.AccessRefreshTokens, error)

	AuthWithAccessToken(ctx context.Context, req *auth.AuthWithAccessTokenRequest) (*auth.AuthWithAccessTokenResponse, error)
}
