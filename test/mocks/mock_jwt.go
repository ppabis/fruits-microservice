package mocks

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// If user is -1, don't set the sub claim
func MockJWT(user int, service string) *jwt.Token {

	claims := jwt.MapClaims{
		"aud": fmt.Sprintf("service:%s", service),
		"exp": time.Now().Add(time.Second * 30).Unix(),
	}

	if user != -1 {
		claims["sub"] = fmt.Sprintf("user:%d", user)
	}

	return jwt.NewWithClaims(jwt.SigningMethodES512, &claims)
}

func MockJWTWithExtras(user int, service string, extras map[string]interface{}) *jwt.Token {

	claims := jwt.MapClaims{
		"sub": fmt.Sprintf("user:%d", user),
		"aud": fmt.Sprintf("service:%s", service),
		"exp": time.Now().Add(time.Second * 30).Unix(),
	}

	for key, value := range extras {
		claims[key] = value
	}

	return jwt.NewWithClaims(jwt.SigningMethodES512, claims)
}
