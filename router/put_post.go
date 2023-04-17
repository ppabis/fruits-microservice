package router

import (
	"fmt"
	"fruits_microservice/auth"
	"fruits_microservice/fruits"
	"net/url"

	"github.com/golang-jwt/jwt/v5"
)

var ErrAuthorization = fmt.Errorf("authorization failed")
var ErrBadRequest = fmt.Errorf("bad request")

func updateFruit(form url.Values, token *jwt.Token) error {
	err := auth.Authorize(form, token)
	if err != nil {
		return ErrAuthorization
	}

	claims := token.Claims.(jwt.MapClaims)
	subject, err := claims.GetSubject()
	if err != nil {
		return ErrBadRequest
	}

	username := claims["username"].(string)

	return fruits.UpdateFruit(subject, username, form.Get("fruit"))
}
