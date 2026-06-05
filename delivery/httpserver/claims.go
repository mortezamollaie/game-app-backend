package httpserver

import (
	"game-app/pkg/constant"
	authservice "game-app/service/authService"

	"github.com/labstack/echo/v4"
)

func GetClaims(c echo.Context) *authservice.Claims {
	return c.Get(constant.AuthMiddlewareContextKey).(*authservice.Claims)
}
