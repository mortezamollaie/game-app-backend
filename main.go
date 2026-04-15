package main

import (
	"game-app/config"
	"game-app/delivery/httpserver"
	"game-app/repository/mysql"
	authservice "game-app/service/authService"
	userservice "game-app/service/userservice"

	"time"
)

const (
	JwtSignKey                 = "jwt_secret_key"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rf"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func main() {

	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8080},
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
		Mysql: mysql.Config{
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Password: "",
			DBName:   "gameapp_db",
		},
	}

	authSvc, userSvc := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc)
	server.Serve()
}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	authSvc := authservice.New(cfg.Auth)

	MysqlRepo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(authSvc, MysqlRepo)
	return authSvc, userSvc
}
