package config

import "time"

const (
	JwtSignKey                 = "jwt_secret_key"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rf"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
	AuthMiddlewareContextKey   = "claims"
)
