package backofficeuserhandler

import (
	authservice "game-app/service/authService"
	"game-app/service/authorizationservice"
	"game-app/service/backofficeuserservice"
)

type Handler struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	authorizationSvc  authorizationservice.Service
	backofficeUserSvc backofficeuserservice.Service
}

func New(
	config authservice.Config,
	authSvc authservice.Service,
	authorizationSvc authorizationservice.Service,
	backofficeUserSvc backofficeuserservice.Service,
) Handler {
	return Handler{
		authConfig:        config,
		authSvc:           authSvc,
		authorizationSvc:  authorizationSvc,
		backofficeUserSvc: backofficeUserSvc,
	}
}
