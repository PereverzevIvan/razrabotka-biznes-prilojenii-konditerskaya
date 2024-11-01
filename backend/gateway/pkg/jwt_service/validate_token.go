package jwt_service

import (
	"errors"

	jwt "github.com/golang-jwt/jwt"
)

const (
	ErrTokenEmpty   = "token is empty"
	ErrTokenInvalid = "token is invalid"
)

func (j *JWTService) ValidateToken(token string) (*jwt.Token, error) {
	if token == "" {
		return nil, errors.New(ErrTokenEmpty)
	}

	jwt_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(j.config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !jwt_token.Valid {
		return nil, errors.New(ErrTokenInvalid)
	}

	return jwt_token, nil
}
