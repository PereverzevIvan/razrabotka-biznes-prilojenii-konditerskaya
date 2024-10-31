package jwt_service

import "time"

type ConfigJWT struct {
	SecretKey              string
	AccessTokenExpiration  time.Duration
	RefreshTokenExpiration time.Duration
}

type iUser interface {
	GetID() int
	GetRole() string
}

// type iToken interface {
// 	GetID() int
// 	GetEmail() string
// 	GetRole() string
// 	GetVersion() int
// }
