package e2e

import (
	"crypto/ecdsa"
	"net/http"

	"github.com/ory/dockertest/v3"
)

var redisContainer *dockertest.Resource
var httpPort int
var goodKey *ecdsa.PrivateKey
var badKey *ecdsa.PrivateKey
var client = &http.Client{}
