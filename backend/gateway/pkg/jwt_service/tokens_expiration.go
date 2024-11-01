package jwt_service

import "time"

func (j *JWTService) AccessTokenExpiration() time.Duration {
	return j.config.AccessTokenExpiration
}

func (j *JWTService) RefreshTokenExpiration() time.Duration {
	return j.config.RefreshTokenExpiration
}
