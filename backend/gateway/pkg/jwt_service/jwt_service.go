package jwt_service

import (
	"errors"
	"time"

	// jwt "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/jwt"
	jwt "github.com/golang-jwt/jwt"
)

const (
	ErrTokenEmpty   = "token is empty"
	ErrTokenInvalid = "token is invalid"
)

type jwtService struct {
	config *ConfigJWT
}

func NewJWTService(jwtConfig *ConfigJWT) jwtService {
	return jwtService{
		config: jwtConfig,
	}
}

func (j jwtService) generateToken(user iUser, expiration time.Duration) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.GetID(),
		"role": user.GetRole(),
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(expiration).Unix(),
		// "nbf":   user.NotBefore,
		// "jti":   user.JTI,
	})

	return token.SignedString([]byte(j.config.SecretKey))
}

func (j jwtService) GenerateAccessAndRefreshTokens(user iUser) (access_token string, refresh_token string, err error) {
	// Генерация токенов
	access_token, err = j.generateToken(user, j.config.AccessTokenExpiration)
	if err != nil {
		return "", "", err
	}

	refresh_token, err = j.generateToken(user, j.config.RefreshTokenExpiration)
	if err != nil {
		return "", "", err
	}

	return access_token, refresh_token, nil
}

func (j jwtService) ValidateToken(token string) (*jwt.Token, error) {
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

func (j jwtService) AccessTokenExpiration() time.Duration {
	return j.config.AccessTokenExpiration
}

func (j jwtService) RefreshTokenExpiration() time.Duration {
	return j.config.RefreshTokenExpiration
}
