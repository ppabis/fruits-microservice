package config

import (
	"crypto/ecdsa"
	"io"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var RedisEndpoint string = "localhost:6379"
var PublicKey *ecdsa.PublicKey = nil

func init() {
	if os.Getenv("REDIS_ENDPOINT") != "" {
		RedisEndpoint = os.Getenv("REDIS_ENDPOINT")
	}

	log.Default().Printf("info: redis endpoint set to %q\n", RedisEndpoint)

	publicKeyPath := "monolith.pem"
	if os.Getenv("PUBLIC_KEY_FILE") != "" {
		if _, err := os.Stat(os.Getenv("PUBLIC_KEY_FILE")); err == nil {
			publicKeyPath = os.Getenv("PUBLIC_KEY_FILE")
		}
	}

	log.Default().Printf("info: public key path set to %q\n", publicKeyPath)

	loadPublicKey(publicKeyPath)

}

func loadPublicKey(path string) {
	publicKeyFile, err := os.Open(path)
	if err != nil {
		log.Default().Printf("failed to open public key: %v\n", err)
		panic(err)
	}

	keyBytes, err := io.ReadAll(publicKeyFile)
	if err != nil {
		log.Default().Printf("failed to read public key: %v\n", err)
		panic(err)
	}

	PublicKey, err = jwt.ParseECPublicKeyFromPEM(keyBytes)
	if err != nil {
		log.Default().Printf("failed to parse public key: %v\n", err)
		panic(err)
	}

	log.Default().Printf("loaded public key, curve %q\n", PublicKey.Curve.Params().Name)
}
