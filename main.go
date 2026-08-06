package main

import (
	"fmt"
	"game-app/adapter/redis"
	"game-app/config"
	"game-app/delivery/httpserver"
	"game-app/repository/migrator"
	"game-app/repository/mysql"
	"game-app/repository/mysql/mysqlaccesscontrol"
	"game-app/repository/mysql/mysqluser"
	"game-app/repository/redis/redismatching"
	authservice "game-app/service/authService"
	"game-app/service/authorizationservice"
	"game-app/service/backofficeuserservice"
	"game-app/service/matchingservice"
	userservice "game-app/service/userservice"
	"game-app/validator/matchingvalidator"
	"game-app/validator/uservalidator"
)

func main() {
	cfg := config.Load("config.yml")
	fmt.Println("cfg: ", cfg)

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

	redisAdapter := redis.New(cfg.Redis)
	matchingRepo := redismatching.New(redisAdapter)
	matchingValidator := matchingvalidator.New()
	matchingSvc := matchingservice.New(cfg.MatchingService, matchingRepo)

	return authSvc, userSvc, uV, backofficeUserSvc, authorizationSvc, matchingSvc, matchingValidator
}
