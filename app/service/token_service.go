package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/midaef/emmet-server/app/models"
	"github.com/midaef/emmet-server/tools/helpers"
	jwt_helper "github.com/midaef/emmet-server/tools/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *Service) CreateToken(ctx context.Context, token *models.Token) error {
	return s.repository.TokenRepository.CreateToken(ctx, token)
}

func (s *Service) GetTokenByAccessToken(ctx context.Context, accessToken string) (*models.Token, error) {
	return s.repository.TokenRepository.GetTokenByAccessToken(ctx, accessToken)
}

func (s *Service) DeleteTokenByAccessToken(ctx context.Context, accessToken string) error {
	return s.repository.TokenRepository.DeleteTokenByAccessToken(ctx, accessToken)
}

func (s *Service) GenerateTokens(ctx context.Context, userID uint64, jwtSecretKey string, expirationTime time.Duration) (*models.Token, error) {
	refreshToken := helpers.GenerateRandomString(64)
	accessToken, err := jwt_helper.CreateJWT([]byte(jwtSecretKey), &jwt_helper.Claims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
		},
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "token creation error: "+err.Error())
	}

	token := &models.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       userID,
		Exp:          time.Now().Add(expirationTime),
	}

	return token, nil
}
