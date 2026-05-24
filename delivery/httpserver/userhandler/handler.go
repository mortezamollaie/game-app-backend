package userhandler

import (
	authservice "game-app/service/authService"
	"game-app/service/userservice"
	"game-app/validator/uservalidator"
)

type Handler struct {
	authConfig    authservice.Config
	authSvc       authservice.Service
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}

func New(config authservice.Config, authSvc authservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator) Handler {
	return Handler{
		authConfig:    config,
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}
