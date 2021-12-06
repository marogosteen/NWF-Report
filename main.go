package main

import (
	"log"

	"example.com/osunnwf.report/app/controllers"
)

func main() {
	err := controllers.StartWebServer()
	if err != nil {
		log.Fatal(err)
	}
}
