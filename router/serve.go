package router

import (
	"fruits_microservice/auth"
	"net/http"
	"strconv"
)

func Serve(port int) error {
	// Serve the web app
	mux := http.NewServeMux()
	mux.HandleFunc("/fruit/", user)
	mux.HandleFunc("/fruit", fruit)
	mux.HandleFunc("/", root)
	return http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

func validateMethod(w http.ResponseWriter, r *http.Request, allowed ...string) bool {
	for _, method := range allowed {
		if r.Method == method {
			return true
		}
	}
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method not allowed"))
	return false
}

func fruit(w http.ResponseWriter, r *http.Request) {
	// handles URL: /fruit with POST and PUT - updates user's fruit
	if !validateMethod(w, r, "POST", "PUT") {
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to parse form"))
		return
	}

	token := auth.Authenticate(w, r)
	if token == nil {
		return
	}

	// TODO: update fruit
}

func user(w http.ResponseWriter, r *http.Request) {
	// handles URL: /fruit/:id - gets single user's fruit
	if !validateMethod(w, r, "GET") {
		return
	}

	getByUser(w, r)

}

func root(w http.ResponseWriter, r *http.Request) {
	// handles URL: / - returns all fruits
	if !validateMethod(w, r, "GET") {
		return
	}

	getAllFruits(w)

}
