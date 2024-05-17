package main

import (
	"os"
	"os/signal"
	"syscall"
	"tools/internals/config"
	"tools/internals/handler"
	"tools/internals/repository"
	"tools/internals/server"
	"tools/internals/service"

	"github.com/sirupsen/logrus"
)

func main() {
	conf, err := config.LoadAndSaveConfig("../config/config.yaml")
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	db, err := repository.NewPostgresDB(conf)
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	defer db.Close()

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)
	server := server.NewServer(handler.InitRoutes(), conf)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	logrus.Printf("Server started listening on port: %s\n", conf.Server.Port)

	select {
	case s := <-interrupt:
		logrus.Printf("Received signal: %s\n", s.String())
	case err = <-server.Notify():
		logrus.Fatalf("Notify: %v", err)
	}

	err = server.Shutdown()
	if err != nil {
		logrus.Fatalf("Received error when shutting down: %v", err)
	}

	logrus.Println("Server stopped")
}
