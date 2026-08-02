package main

import (
	"game-app/config"
	"game-app/delivery/httpserver"
	"game-app/repository/migrator"
	"game-app/repository/mysql"
	"game-app/repository/mysql/mysqlaccesscontrol"
	"game-app/repository/mysql/mysqluser"
	authservice "game-app/service/authService"
	"game-app/service/authorizationservice"
	"game-app/service/backofficeuserservice"
	"game-app/service/matchingservice"
	userservice "game-app/service/userservice"
	"game-app/validator/matchingvalidator"
	"game-app/validator/uservalidator"
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
	config.Load()

	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8088},
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

	// TODO: add command for migrations
	mgr := migrator.New(cfg.Mysql, "mysql")
	mgr.Up()

	// TODO: add struct and add this returned item to struct fields
	authSvc, userSvc, userValidator, backofficeSvc, authorizationSvc, matchingSvc, matchingValidator := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc, userValidator, backofficeSvc, authorizationSvc, matchingSvc, matchingValidator)
	server.Serve()
}

func setupServices(cfg config.Config) (
	authservice.Service,
	userservice.Service,
	uservalidator.Validator,
	backofficeuserservice.Service,
	authorizationservice.Service,
	matchingservice.Service,
	matchingvalidator.Validator) {
	authSvc := authservice.New(cfg.Auth)

	MysqlRepo := mysql.New(cfg.Mysql)

	userMysql := mysqluser.New(MysqlRepo)
	aclMysql := mysqlaccesscontrol.New(MysqlRepo)
	userSvc := userservice.New(authSvc, userMysql)

	backofficeUserSvc := backofficeuserservice.New()
	authorizationSvc := authorizationservice.New(aclMysql)

	uV := uservalidator.New(userMysql)

	matchingValidator := matchingvalidator.New()
	matchingSvc := matchingservice.New(cfg.MatchingService)

	return authSvc, userSvc, uV, backofficeUserSvc, authorizationSvc, matchingSvc, matchingValidator
}
