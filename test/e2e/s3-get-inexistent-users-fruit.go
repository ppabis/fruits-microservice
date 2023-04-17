package e2e

import (
	"net/http"
	"strconv"
	"testing"
)

/*
Scenario 3: Get inexistent user's fruit
*/
func Get_Inexistent_Users_Fruit(t *testing.T) {
	resp, err := client.Get("http://localhost:" + strconv.Itoa(httpPort) + "/fruit/10")
	if err != nil {
		t.Errorf("Error while getting one user's fruit: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %v, got %v", http.StatusNotFound, resp.StatusCode)
	}
}
