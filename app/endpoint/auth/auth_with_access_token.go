package auth

import (
	"context"
	"github.com/midaef/emmet-server/extra/auth"
	jwt_helper "github.com/midaef/emmet-server/tools/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (a *AuthEndpoint) AuthWithAccessToken(ctx context.Context, req *auth.AuthWithAccessTokenRequest) (*auth.AuthWithAccessTokenResponse, error) {
	token, err := jwt_helper.ParseJWT([]byte(a.config.JWT.SecretKey), req.AccessToken)
	if err != nil {
		return nil, err
	}

	if token.ExpiresAt < time.Now().Unix() {
		return nil, status.Error(codes.Unauthenticated, "token has expired")
	}

	return &auth.AuthWithAccessTokenResponse{}, nil
}
