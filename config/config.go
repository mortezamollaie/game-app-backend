package config

import (
	"game-app/repository/mysql"
	authservice "game-app/service/authService"
)

type HTTPServer struct {
	Port int
}

type Config struct {
	HTTPServer HTTPServer
	Auth       authservice.Config
	Mysql      mysql.Config
}
