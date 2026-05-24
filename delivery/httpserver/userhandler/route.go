package userhandler

import (
	"game-app/delivery/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userG := e.Group("/users")

	userG.POST("/login", h.userLogin)
	userG.GET("/profile", h.userProfile, middleware.Auth(h.authSvc, h.authConfig))
	userG.POST("/register", h.userRegister)
}
