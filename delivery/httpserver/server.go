package httpserver

import (
	"fmt"
	"game-app/config"
	"game-app/delivery/httpserver/backofficeuserhandler"
	"game-app/delivery/httpserver/matchinghandler"
	"game-app/delivery/httpserver/userhandler"
	authservice "game-app/service/authService"
	"game-app/service/authorizationservice"
	"game-app/service/backofficeuserservice"
	"game-app/service/matchingservice"
	userservice "game-app/service/userservice"
	"game-app/validator/matchingvalidator"
	"game-app/validator/uservalidator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config                config.Config
	userHandler           userhandler.Handler
	backofficeUserHandler backofficeuserhandler.Handler
	matchingHandler       matchinghandler.Handler
}

func New(
	config config.Config,
	authSvc authservice.Service,
	userSvc userservice.Service,
	userValidator uservalidator.Validator,
	backofficeUserSvc backofficeuserservice.Service,
	authorizationSvc authorizationservice.Service,
	matchingSvc matchingservice.Service,
	matchingValidator matchingvalidator.Validator,
) Server {
	return Server{
		config:                config,
		userHandler:           userhandler.New(config.Auth, authSvc, userSvc, userValidator),
		backofficeUserHandler: backofficeuserhandler.New(config.Auth, authSvc, authorizationSvc, backofficeUserSvc),
		matchingHandler:       matchinghandler.New(config.Auth, authSvc, matchingSvc, matchingValidator),
	}
}

func (s Server) Serve() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check", s.healthCheck)

	s.userHandler.SetUserRoutes(e)
	s.backofficeUserHandler.SetUserListRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))
}
