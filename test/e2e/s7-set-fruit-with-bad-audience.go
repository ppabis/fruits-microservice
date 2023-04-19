package e2e

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Set_Fruit_With_Bad_Audience(t *testing.T) {
	tok, err := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.MapClaims{
		"username": "johnathan",
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
		"aud":      "service:wrong",
		"sub":      "user:2",
		"super":    false,
	}).SignedString(goodKey)

	if err != nil {
		t.Errorf("Error generating token: %v", err)
	}

	form := url.Values{}
	form.Set("fruit", "watermelon")

	req, err := http.NewRequest("PUT", "http://localhost:"+strconv.Itoa(httpPort)+"/fruit", strings.NewReader(form.Encode()))
	if err != nil {
		t.Errorf("Error while creating request: %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	req.Header.Add("X-Auth-Token", tok)

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error while setting one user's fruit: %v", err)
	}

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %v", resp.StatusCode)
	}
}
