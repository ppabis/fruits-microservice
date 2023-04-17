package router

import (
	"fruits_microservice/auth"
	"net/http"
	"strconv"
)

var Server *http.Server

func Serve(port int) error {
	// Serve the web app
	mux := http.NewServeMux()
	mux.HandleFunc("/fruit/", user)
	mux.HandleFunc("/fruit", fruit)
	mux.HandleFunc("/", root)
	Server = &http.Server{Addr: ":" + strconv.Itoa(port), Handler: mux}
	return Server.ListenAndServe()
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

	err = updateFruit(r.Form, token)
	if err != nil {
		if err == ErrAuthorization {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Authorization failed"))
		} else if err == ErrBadRequest {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
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
