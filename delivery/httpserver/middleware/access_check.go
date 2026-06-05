package middleware

import (
	"game-app/pkg/claim"
	"game-app/pkg/errmsg"
	"game-app/service/authorizationservice"
	"net/http"

	"github.com/labstack/echo/v4"
)

// closure or higher order function => function that return function

func AccessCheck(service authorizationservice.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := claim.GetClaimsFromEchoContext(c)
			isAllowed, err := service.CheckAccess(claims.UserID, claims.Role)
			if err != nil {
				// TODO - log unexpected error
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": errmsg.SomethingWentWrong,
				})
			}

			if !isAllowed {
				return c.JSON(http.StatusForbidden, echo.Map{
					"message": errmsg.ErrorMsgAccessDenied,
				})
			}
			return next(c)
		}
	}
}
