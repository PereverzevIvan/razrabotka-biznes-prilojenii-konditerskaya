package jwt_service

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/jwt_service"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt"
)

type JWTService struct {
	jwtService *jwt_service.JWTService
}

func NewJWTService(jwtConfig *configs.JWTConfig) *JWTService {
	return &JWTService{
		jwtService: jwt_service.NewJWTService(
			&jwt_service.ConfigJWT{
				SecretKey:              jwtConfig.SecretKey,
				AccessTokenExpiration:  jwtConfig.AccessTokenDuration,
				RefreshTokenExpiration: jwtConfig.RefreshTokenDuration,
			}),
	}
}

func (service *JWTService) ValidateToken(token string) (*jwt.Token, error) {
	return service.jwtService.ValidateToken(token)
}

func (service *JWTService) parseRoleFromToken(token *jwt.Token) (string, error) {
	return service.jwtService.GetRole(token)
}

func (service *JWTService) parseUserIDFromToken(token *jwt.Token) (int, error) {
	return service.jwtService.GetUserID(token)
}

func (service *JWTService) GetRoleFromAccessTokenCookie(ctx fiber.Ctx) (string, error) {
	access_token, err := service.jwtService.ValidateToken(ctx.Cookies("access_token"))
	if err != nil {
		return "", err
	}

	return service.parseRoleFromToken(access_token)
}

func (service *JWTService) GetUserIDFromAccessTokenCookie(ctx fiber.Ctx) (int, error) {
	access_token, err := service.jwtService.ValidateToken(ctx.Cookies("access_token"))
	if err != nil {
		return 0, err
	}

	return service.parseUserIDFromToken(access_token)
}

func (service *JWTService) GetUserIDFromRefreshTokenCookie(ctx fiber.Ctx) (int, error) {
	refresh_token, err := service.jwtService.ValidateToken(ctx.Cookies("refresh_token"))
	if err != nil {
		return 0, err
	}

	return service.parseUserIDFromToken(refresh_token)
}
