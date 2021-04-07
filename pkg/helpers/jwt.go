package helpers

import (
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"time"
)

// JWTManager ...
type JWTManager interface {
	CreateAccessToken(claims *Claims) (string, error)
	ParseJWT(accessToken string) (*Claims, error)
	CreateRefreshToken() string
	IsCorrectJWT(accessToken string) (bool, error)
}

type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

type JWT struct {
	secretKey string
}

func NewJWT(secretKey string) (*JWT, error) {
	if secretKey == "" {
		return nil, errors.New("jwt: private key is empty")
	}

	return &JWT{
		secretKey: secretKey,
	}, nil
}

func (j *JWT) CreateAccessToken(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}

func (j *JWT) ParseJWT(accessToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected token signing method")
			}

			return []byte(j.secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*Claims)

	return claims, nil
}

func (j *JWT) IsCorrectJWT(accessToken string) (bool, error) {
	claims, _ := j.ParseJWT(accessToken)

	if claims.ExpiresAt < time.Now().Unix() {
		return false, errors.New("Token lifetime expired")
	}

	token, _ := j.CreateAccessToken(claims)
	if token == accessToken {
		return true, nil
	}

	return false, nil
}

func (j *JWT) CreateRefreshToken() string {
	token := securecookie.GenerateRandomKey(64)
	strToken := base64.StdEncoding.EncodeToString(token)

	return strToken
}
