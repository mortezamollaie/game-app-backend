package matchinghandler

import (
	"game-app/delivery/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userG := e.Group("/matching")

	userG.POST("/add-to-waiting-list", h.addToWaitingList, middleware.Auth(h.authSvc, h.authConfig))
}
