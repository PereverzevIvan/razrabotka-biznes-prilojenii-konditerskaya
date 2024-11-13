package controllers

import (
	"time"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"
	params_auth "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models/params/auth"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt"
)

type IUserService interface {
	Create(params *params_auth.RegisterParams) (*models.User, error)
	GetByID(user_id int) (*models.User, error)
	GetByLogin(email string) (*models.User, error)
	IsPasswordCorrect(user *models.User, password string) bool
	// GetUser(userId int) (string, error)
}

type IJWTService interface {
	GenerateAccessAndRefreshTokens(user *models.User) (access_token string, refresh_token string, err error)
	ValidateToken(token string) (*jwt.Token, error)

	AccessTokenExpiration() time.Duration
	RefreshTokenExpiration() time.Duration

	GetRoleFromAccessTokenCookie(ctx fiber.Ctx) (string, error)
	GetUserIDFromAccessTokenCookie(ctx fiber.Ctx) (int, error)

	GetUserIDFromRefreshTokenCookie(ctx fiber.Ctx) (int, error)
}

type IAuthMiddleware interface {
	Authorized(ctx fiber.Ctx) error
	// Admin(ctx fiber.Ctx) error
}
