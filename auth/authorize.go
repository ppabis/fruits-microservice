package auth

import (
	"fmt"
	"fruits_microservice/fruits"
	"net/url"

	"github.com/golang-jwt/jwt/v5"
)

func Authorize(form url.Values, token *jwt.Token) error {
	claims := token.Claims.(jwt.MapClaims)

	// Check if the user tries to set a special fruit
	// and if the user is allowed to do so
	super := claims["super"].(bool)
	if fruits.IsFruitSpecial(form.Get("fruit")) && !super {
		return fmt.Errorf("this user is not allowed to set special fruit")
	}

	return nil
}
