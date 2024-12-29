package main

import (
	"fmt"
	"klambri-backend/internal/configuration"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := configuration.ReadConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	if err := r.Run(fmt.Sprintf(":%v", config.Server.Port)); err != nil {
		panic(err)
	}
}
