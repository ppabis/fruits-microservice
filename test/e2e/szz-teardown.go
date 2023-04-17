package e2e

import (
	"context"
	"fruits_microservice/router"
	"fruits_microservice/test/integration"
	"log"
)

func Teardown() {
	err := integration.DestroyContainer(redisContainer)
	log.Default().Println("Redis container destroyed")
	err2 := router.Server.Shutdown(context.TODO())
	log.Default().Println("Server stopped")

	if err != nil || err2 != nil {
		log.Default().Printf("Error while tearing down: %v, %v", err, err2)
	}
}
