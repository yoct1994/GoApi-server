package main

import (
	"github.com/GoAPI-server/server"
	"github.com/savsgio/go-logger/v2"
)

func main() {
	server.Start(7000)

	logger.Info("Starting server")
}
