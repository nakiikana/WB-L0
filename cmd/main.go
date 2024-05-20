package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"tools/internals/config"
	"tools/internals/handler"
	"tools/internals/repository"
	"tools/internals/server"
	"tools/internals/service"
	st "tools/internals/stan"
	publicher "tools/internals/stan/stan-pub"
	subscriber "tools/internals/stan/stan-sub"

	"github.com/nats-io/stan.go"
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

	//creating nats connection
	nc, err := st.Connect()
	if err != nil {
		logrus.Fatalf("%v\n", err)
	}
	defer nc.CloseConnection()

	//creating stan connection
	log.Println(nc.URL)
	sc, err := stan.Connect("test", "subscriber", stan.NatsURL(nc.URL))
	if err != nil {
		logrus.Fatalf("%v\n", err)
	}
	logrus.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", nc.URL, "test", "subscriber")
	pub := publicher.NewPublisher(sc)
	_, err = subscriber.NewSubscriber(sc, "order")
	enough := make(chan bool, 1)
	if err != nil {
		logrus.Fatalf("%v\n", err)
	}

	go func() {
		if err := pub.SendWithTimeout(enough); err != nil {
			logrus.Fatalf("%v", err)
		}
	}()

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
