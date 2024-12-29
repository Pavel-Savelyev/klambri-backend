package main

import (
	"fmt"
	"klambri-backend/internal/configuration"
	natsclient "klambri-backend/internal/nats-client"
	"klambri-backend/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := configuration.ReadConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if err := natsclient.Init(config.NatsServer.Url); err != nil {
		log.Fatal("Error ", err)
	}
	defer natsclient.Close()

	r := gin.Default()

	routes.SetupRoutes(r)

	if err := r.Run(fmt.Sprintf(":%v", config.Server.Port)); err != nil {
		panic(err)
	}
}
