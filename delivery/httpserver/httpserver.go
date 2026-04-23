package httpserver

import (
	"fmt"
	"game-app/config"
	authservice "game-app/service/authService"
	userservice "game-app/service/userservice"
	"game-app/validator/uservalidator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config        config.Config
	authSvc       authservice.Service
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator) Server {
	return Server{
		config:        config,
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}

func (s Server) Serve() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check", s.healthCheck)

	userG := e.Group("/users")

	userG.POST("/login", s.userLogin)
	userG.GET("/profile", s.userProfile)
	userG.POST("/register", s.userRegister)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))
}
