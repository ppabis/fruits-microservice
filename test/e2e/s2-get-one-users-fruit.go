package e2e

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"
)

/*
Scenario 2: Get one user's fruit
*/
func Get_One_Users_Fruit(t *testing.T) {
	resp, err := client.Get("http://localhost:" + strconv.Itoa(httpPort) + "/fruit/2")
	if err != nil {
		t.Errorf("Error while getting one user's fruit: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
	}

	fruit := struct {
		Username string `json:"username"`
		Fruit    string `json:"fruit"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&fruit)
	if err != nil {
		t.Errorf("Error while decoding response: %v", err)
	}

	if fruit.Username != "johnathan" {
		t.Errorf("Expected username to be johnathan, got %s", fruit.Username)
	}

	if fruit.Fruit != "apple" {
		t.Errorf("Expected fruit to be apple, got %s", fruit.Fruit)
	}
}
