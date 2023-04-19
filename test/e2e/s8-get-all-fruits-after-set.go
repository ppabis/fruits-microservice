package e2e

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"
)

/*
Scenario 8: Get all fruits from the database after changing one
*/
func Get_All_Fruits_After_Set(t *testing.T) {
	validMap := map[string]string{
		"john":      "kiwi",
		"johnathan": "banana",
		"damian":    "apple",
		"alexis":    "pineapple",
	}

	resp, err := client.Get("http://localhost:" + strconv.Itoa(httpPort) + "/")
	if err != nil {
		t.Errorf("Error while getting all fruits: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
	}

	fruitsList := make([]struct {
		Username string `json:"username"`
		Fruit    string `json:"fruit"`
	}, 0, 512)

	err = json.NewDecoder(resp.Body).Decode(&fruitsList)

	if err != nil {
		t.Errorf("Error while decoding response: %v", err)
	}

	for tuple := range fruitsList {
		if _, ok := validMap[fruitsList[tuple].Username]; !ok {
			t.Errorf("Expected %s to be in the map", fruitsList[tuple].Username)
		}

		if validMap[fruitsList[tuple].Username] != fruitsList[tuple].Fruit {
			t.Errorf("Expected %s's fruit to be %s, got %s", fruitsList[tuple].Username, validMap[fruitsList[tuple].Username], fruitsList[tuple].Fruit)
		}
	}
}
