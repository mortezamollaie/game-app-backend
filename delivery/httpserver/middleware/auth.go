package middleware

import (
	"game-app/pkg/constant"
	authservice "game-app/service/authService"

	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// closure or higher order function => function that return function

func Auth(service authservice.Service, config authservice.Config) echo.MiddlewareFunc {
	return mw.WithConfig(mw.Config{
		ContextKey: constant.AuthMiddlewareContextKey,
		SigningKey: []byte(config.SignKey),
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claims, err := service.ParseToken(auth)
			if err != nil {
				return nil, err
			}
			return claims, nil
		},
	})
}
