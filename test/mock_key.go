package test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fruits_microservice/config"
)

func MockKeyPair() *ecdsa.PrivateKey {
	key, _ := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	config.PublicKey = &key.PublicKey
	return key
}
