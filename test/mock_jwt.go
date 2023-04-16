package test

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func MockJWT(user int, service string) *jwt.Token {

	claims := jwt.MapClaims{
		"sub": fmt.Sprintf("user:%d", user),
		"aud": fmt.Sprintf("service:%s", service),
		"exp": time.Now().Add(time.Second * 30).Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodES512, claims)
}
