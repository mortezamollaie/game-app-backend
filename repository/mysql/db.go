package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Username string
	Password string
	Port     int
	Host     string
	DBName   string
}

type MySQLDB struct {
	config Config
	db     *sql.DB
}

func New(config Config) *MySQLDB {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.Username, config.Password, config.Host, config.Port, config.DBName),
	)
	if err != nil {
		panic(fmt.Errorf("can't open database %v", err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{
		config: config,
		db:     db}

}
