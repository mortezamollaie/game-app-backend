package matchinghandler

import (
	authservice "game-app/service/authService"
	"game-app/service/matchingservice"
	"game-app/validator/matchingvalidator"
)

type Handler struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	matchingSvc       matchingservice.Service
	matchingValidator matchingvalidator.Validator
}

func New(config authservice.Config, authSvc authservice.Service, matchingSvc matchingservice.Service, matchingValidator matchingvalidator.Validator) Handler {
	return Handler{
		authConfig:        config,
		authSvc:           authSvc,
		matchingSvc:       matchingSvc,
		matchingValidator: matchingValidator,
	}
}
