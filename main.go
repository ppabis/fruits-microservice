package main

import (
	"fruits_microservice/router"
	"log"
)

func main() {
	err := router.Serve(8081)
	if err != nil {
		log.Default().Printf("Shutting down server: %v\n", err)
		panic(err)
	}
}
