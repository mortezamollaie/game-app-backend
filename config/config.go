package config

import (
	"game-app/repository/mysql"
	authservice "game-app/service/authService"
)

type HTTPServer struct {
	Port `koanf:"port"`
}

type Config struct {
	HTTPServer HTTPServer         `koanf:"http_server"`
	Auth       authservice.Config `koanf:"auth"`
	Mysql      mysql.Config       `koanf:"mysql"`
}
