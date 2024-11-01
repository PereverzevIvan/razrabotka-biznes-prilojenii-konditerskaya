package jwt_service

import (
	"errors"

	jwt "github.com/golang-jwt/jwt"
)

type JWTService struct {
	config *ConfigJWT
}

func NewJWTService(jwtConfig *ConfigJWT) *JWTService {
	return &JWTService{
		config: jwtConfig,
	}
}

func (j *JWTService) GetRole(token *jwt.Token) (string, error) {
	if token == nil {
		return "", errors.New("token is nil")
	}

	role_claims := token.Claims.(jwt.MapClaims)["role"]
	if role_claims == nil {
		return "", errors.New("invalid role in token")
	}

	role := role_claims.(string)
	return role, nil
}
func (j *JWTService) GetUserID(token *jwt.Token) (int, error) {
	if token == nil {
		return 0, errors.New("token is nil")
	}

	user_id_claims := token.Claims.(jwt.MapClaims)["id"]
	if user_id_claims == nil {
		return 0, errors.New("invalid user id in token")
	}

	user_id := int(user_id_claims.(float64))
	return user_id, nil
}
