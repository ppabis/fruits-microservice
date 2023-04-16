package auth

import (
	"fruits_microservice/test"
	"net/url"
	"testing"
)

func Test_SpecialFruitSuperUser(t *testing.T) {

	token := test.MockJWTWithExtras(1, "fruits", map[string]interface{}{
		"super":    true,
		"username": "Henry",
	})

	form := make(url.Values)
	form.Set("fruit", "kiwi")

	err := Authorize(form, token)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_SpecialFruitNormalUser(t *testing.T) {
	token := test.MockJWTWithExtras(1, "fruits", map[string]interface{}{
		"super":    false,
		"username": "Henry",
	})

	form := make(url.Values)
	form.Set("fruit", "kiwi")

	err := Authorize(form, token)

	if err == nil {
		t.Error("This should throw an error")
	}
}

func Test_NormalFruitNormalUser(t *testing.T) {
	token := test.MockJWTWithExtras(1, "fruits", map[string]interface{}{
		"super":    false,
		"username": "Henry",
	})

	form := make(url.Values)
	form.Set("fruit", "banana")

	err := Authorize(form, token)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_NormalFruitSuperUser(t *testing.T) {
	token := test.MockJWTWithExtras(1, "fruits", map[string]interface{}{
		"super":    true,
		"username": "Henry",
	})

	form := make(url.Values)
	form.Set("fruit", "banana")

	err := Authorize(form, token)

	if err != nil {
		t.Fatal(err)
	}
}
