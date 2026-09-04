package main

import (
	"fmt"
	"game-app/config"
	"game-app/repository/migrator"
	"game-app/scheduler"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg := config.Load("config.yml")
	fmt.Println("cfg: ", cfg)

	// TODO: add command for migrations
	mgr := migrator.New(cfg.Mysql, "mysql")
	mgr.Up()

	done := make(chan bool)

	go func() {
		sch := scheduler.New()
		sch.Start(done)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("received interrupt signal. shutting down gracefully...")
	done <- true
	time.Sleep(cfg.Application.GracefulShutdownTimeout)
}
