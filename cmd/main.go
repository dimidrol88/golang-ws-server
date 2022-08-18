package main

import (
	"github.com/dimidrol88/golang-ws-server/internal/handler"
	"github.com/dimidrol88/golang-ws-server/internal/ws"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
}

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		logrus.Fatalf("Error - Port is not found!")
	}
	logrus.Infof("Started on port: %s", port)

	var (
		server         = ws.NewServer("localhost", port)
		requestHandler = handler.NewHandler()
	)

	if err := server.Run(requestHandler.ConfigureRouter()); err != nil {
		logrus.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
