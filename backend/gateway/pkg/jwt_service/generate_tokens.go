package jwt_service

import (
	"time"

	jwt "github.com/golang-jwt/jwt"
)

func (j *JWTService) generateToken(user IUser, expiration time.Duration) (string, error) {

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

func (j *JWTService) GenerateAccessAndRefreshTokens(user IUser) (access_token string, refresh_token string, err error) {
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
