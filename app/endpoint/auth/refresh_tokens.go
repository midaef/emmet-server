package auth

import (
	"context"
	"github.com/midaef/emmet-server/extra/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *AuthEndpoint) RefreshTokens(ctx context.Context, req *auth.AccessRefreshTokens) (*auth.AccessRefreshTokens, error) {
	token, err := a.services.TokenService.GetTokenByAccessToken(ctx, req.GetAccessToken())
	if err != nil {
		return nil, err
	}

	if !(token.RefreshToken == req.GetRefreshToken() && token.AccessToken == req.GetAccessToken()) {
		return nil, status.Error(codes.Unauthenticated, "invalid tokens")
	}

	tokenAuth, err := a.services.TokenService.GenerateTokens(ctx, token.UserID, a.config.JWT.SecretKey, a.config.JWT.ExpirationTime)
	if err != nil {
		return nil, status.Error(codes.Internal, "token creation error")
	}

	err = a.services.TokenService.DeleteTokenByAccessToken(ctx, req.GetAccessToken())
	if err != nil {
		return nil, status.Error(codes.Internal, "error deleting token")
	}

	err = a.services.TokenService.CreateToken(ctx, tokenAuth)
	if err != nil {
		return nil, status.Error(codes.Internal, "token creation error")
	}

	return &auth.AccessRefreshTokens{
		AccessToken:  tokenAuth.AccessToken,
		RefreshToken: tokenAuth.RefreshToken,
	}, nil
}
