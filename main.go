package main

import "fruits_microservice/router"

func main() {
	err := router.Serve(8081)
	if err != nil {
		panic(err)
	}
}
