package middlewares_auth

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/controllers"

type authMiddleware struct {
	jwtService controllers.IJWTService
}

func NewAuthMiddleware(
	jwtService controllers.IJWTService,
) *authMiddleware {
	return &authMiddleware{
		jwtService: jwtService,
	}
}
