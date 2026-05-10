package userhandler

import (
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userG := e.Group("/users")

	userG.POST("/login", h.userLogin)
	userG.GET("/profile", h.userProfile, mw.WithConfig(mw.Config{
		ContextKey: "user",
		SigningKey: h.authSignKey,
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claims, err := h.authSvc.ParseToken(auth)
			if err != nil {
				return nil, err
			}
			return claims, nil
		},
	}))
	userG.POST("/register", h.userRegister)
}
