package router

import (
	"encoding/json"
	"fruits_microservice/fruits"
	"net/http"
	"strconv"
	"strings"
)

func getByUser(w http.ResponseWriter, r *http.Request) {
	// Extract string_id from the URL /fruit/:string_id
	string_id := strings.TrimPrefix(r.URL.Path, "/") // Because path is not guaranteed to start or not with a /
	string_id = strings.TrimPrefix(string_id, "fruit/")
	id, err := strconv.Atoi(string_id)
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}

	username, fruit, err := fruits.GetFruit(id)

	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		if err == fruits.ErrKeyNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Fruit of this user not found"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error occurred while getting fruit"))
		}
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"username": username,
		"fruit":    fruit,
	})
}

type userAndFruit struct {
	Username string `json:"username"`
	Fruit    string `json:"fruit"`
}

func getAllFruits(w http.ResponseWriter) {
	fruits, err := fruits.GetFruits()
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error occurred while getting fruits"))
		return
	}

	usersFruits := make([]userAndFruit, len(fruits))
	i := 0
	for f := range fruits {
		usersFruits[i] = userAndFruit{
			Username: f,
			Fruit:    fruits[f],
		}
		i++
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usersFruits)
}
