package userhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userG := e.Group("/users")

	userG.POST("/login", h.userLogin)
	userG.GET("/profile", h.userProfile)
	userG.POST("/register", h.userRegister)
}
