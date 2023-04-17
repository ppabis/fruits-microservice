package e2e

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	cr "crypto/rand"
	"fruits_microservice/config"
	"fruits_microservice/router"
	"fruits_microservice/test/integration"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// Setup
// Start redis container, Start HTTP server,
// Generate key for signing JWT tokens and a
// bad key to create bad signature
func Setup() error {
	var port int
	var err error
	redisContainer, port, err = integration.RedisWithTestData()
	if err != nil {
		return err
	}

	config.RedisEndpoint = "localhost:" + strconv.Itoa(port)

	httpPort = 58080 + rand.Intn(200)
	serveChan := make(chan error)
	go func() {
		serveChan <- router.Serve(httpPort)
	}()

	log.Default().Println("Waiting for server to start...")
	time.Sleep(time.Duration(100) * time.Millisecond)
	select {
	case err := <-serveChan:
		return err
	default:
		log.Default().Println("Server started")
	}

	goodKey, err = ecdsa.GenerateKey(elliptic.P521(), cr.Reader)
	if err != nil {
		return err
	}

	badKey, err = ecdsa.GenerateKey(elliptic.P521(), cr.Reader)
	if err != nil {
		return err
	}

	config.PublicKey = &goodKey.PublicKey

	return nil
}
