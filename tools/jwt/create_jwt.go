package jwt_helper

import (
	"github.com/dgrijalva/jwt-go"
)

var (
	signingMethod = jwt.SigningMethodHS256
)

type Claims struct {
	ID uint64
	jwt.StandardClaims
}

func CreateJWT(secretKey []byte, claims *Claims) (string, error) {
	token := jwt.NewWithClaims(signingMethod, claims)

	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseJWT(secretKey []byte, tokenByte string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenByte, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*Claims)

	return claims, nil
}
