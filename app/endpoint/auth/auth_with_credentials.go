package auth

import (
	"context"
	"github.com/midaef/emmet-server/app/models"
	"github.com/midaef/emmet-server/extra/auth"
	"github.com/midaef/emmet-server/tools/helpers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *AuthEndpoint) AuthWithCredentials(ctx context.Context, req *auth.AuthWithCredentialsRequest) (*auth.AccessRefreshTokens, error) {
	cred := &models.Credentials{
		Login:    req.GetLogin(),
		Password: req.GetPassword(),
	}

	if err := cred.Validate(); err != nil {
		return nil, err
	}

	if !a.services.UserService.IsExistByLogin(ctx, req.GetLogin()) {
		return nil, status.Error(codes.NotFound, "login does not exist")
	}

	hashPassword, err := helpers.NewMD5Hash(req.GetPassword())
	if err != nil {
		return nil, err
	}

	cred.Password = hashPassword

	id, err := a.services.UserService.GetUserByCredentials(ctx, cred)
	if err != nil {
		return nil, err
	}

	token, err := a.services.TokenService.GenerateTokens(ctx, id, a.config.JWT.SecretKey, a.config.JWT.ExpirationTime)
	if err != nil {
		return nil, err
	}

	if err = a.services.TokenService.CreateToken(ctx, token); err != nil {
		return nil, status.Error(codes.Internal, "token creation error")
	}

	return &auth.AccessRefreshTokens{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
