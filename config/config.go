package config

import (
	"crypto/ecdsa"
	"io"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var RedisEndpoint string = "localhost:6379"
var PublicKey *ecdsa.PublicKey = nil

func init() {
	if os.Getenv("REDIS_ENDPOINT") != "" {
		RedisEndpoint = os.Getenv("REDIS_ENDPOINT")
	}

	publicKeyPath := "monolith.pem"
	if os.Getenv("PUBLIC_KEY_FILE") != "" {
		if _, err := os.Stat(os.Getenv("PUBLIC_KEY_FILE")); err == nil {
			publicKeyPath = os.Getenv("PUBLIC_KEY_FILE")
		}
	}

	loadPublicKey(publicKeyPath)

}

func loadPublicKey(path string) {
	publicKeyFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	keyBytes, err := io.ReadAll(publicKeyFile)
	if err != nil {
		panic(err)
	}

	PublicKey, err = jwt.ParseECPublicKeyFromPEM(keyBytes)
	if err != nil {
		panic(err)
	}
}
