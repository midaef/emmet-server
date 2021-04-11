package service

import (
	"github.com/midaef/emmet-server/pkg/helpers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Token struct {
	hasher          *helpers.Md5
	tokenManager    *helpers.JWT
}

func NewTokenService(hasher *helpers.Md5, tokenManager *helpers.JWT) *Token {
	return &Token{
		hasher:          hasher,
		tokenManager:    tokenManager,
	}
}

func (s *Token) CheckAccessToken(accessToken string) (*helpers.Claims, error) {
	claims, err := s.tokenManager.ParseJWT(accessToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Token incorrect")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, status.Error(codes.Unauthenticated, "Token lifetime expired")
	}

	token, err := s.tokenManager.CreateAccessToken(claims)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Token incorrect")
	}

	if token != accessToken {
		return nil, status.Error(codes.Unauthenticated, "Token incorrect")
	}



	return claims, nil
}