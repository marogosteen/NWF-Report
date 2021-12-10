package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/nwf-report/app/controllers"
	"github.com/nwf-report/config"
)

func main() {
	cfg := config.Config
	router := gin.Default()
	err := controllers.StartWebServer(router, cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}

type Message struct {
	Name string
	Body string
	Time int64
}
