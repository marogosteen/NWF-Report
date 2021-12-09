package main

import (
	"log"

	"github.com/nwf-report/app/controllers"
)

func main() {
	err := controllers.StartWebServer()
	if err != nil {
		log.Fatal(err)
	}
}
