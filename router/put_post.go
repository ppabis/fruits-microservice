package router

import (
	"fmt"
	"fruits_microservice/auth"
	"fruits_microservice/fruits"
	"log"
	"net/url"
	"strings"

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
	if err != nil || subject == "" || !strings.HasPrefix(subject, "user:") {
		log.Default().Printf("bad request, subject %q is invalid\n", subject)
		return ErrBadRequest
	}

	username := claims["username"].(string)

	log.Default().Printf("trying to set fruit %q for %q\n", form.Get("fruit"), subject)

	return fruits.UpdateFruit(subject, username, form.Get("fruit"))
}
