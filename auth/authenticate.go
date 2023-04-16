package auth

import (
	"fruits_microservice/config"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var allowedJwtMethods = []string{jwt.SigningMethodES512.Name}

func Authenticate(w http.ResponseWriter, r *http.Request) *jwt.Token {
	token := r.Header.Get("X-Auth-Token")
	if token == "" {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return nil
	}

	tok, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return config.PublicKey, nil
	}, jwt.WithAudience("service:fruits"), jwt.WithValidMethods(allowedJwtMethods))

	if err != nil || !tok.Valid {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return nil
	}

	return tok
}
