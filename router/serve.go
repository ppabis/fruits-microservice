package router

import (
	"fruits_microservice/auth"
	"fruits_microservice/config"
	"log"
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
	log.Default().Printf("Starting server on port %d\n", port)
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
	log.Default().Printf("[405] [%s] %q not allowed: %q\n", r.RemoteAddr, r.URL.Path, r.Method)
	return false
}

func corsPreflight(w http.ResponseWriter, r *http.Request) bool {
	// handles CORS preflight requests
	w.Header().Add("Access-Control-Allow-Origin", config.Origin)

	if r.Method == "OPTIONS" {
		w.Header().Add("Access-Control-Allow-Methods", "POST, PUT, GET")
		w.Header().Add("Access-Control-Allow-Headers", "Authorization, Content-Type, X-Auth-Token")
		w.WriteHeader(http.StatusOK)
		log.Default().Printf("[200] [%s] %q CORS preflight\n", r.RemoteAddr, r.URL.Path)
		return true
	}
	return false
}

func fruit(w http.ResponseWriter, r *http.Request) {
	// handles URL: /fruit with POST and PUT - updates user's fruit
	if corsPreflight(w, r) {
		return
	}

	if !validateMethod(w, r, "POST", "PUT") {
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to parse form"))
		log.Default().Printf("[400] [%s] %q failed to parse form\n", r.RemoteAddr, r.URL.Path)
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
			log.Default().Printf("[403] [%s] %q forbidden\n", r.RemoteAddr, r.URL.Path)
		} else if err == ErrBadRequest {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))
			log.Default().Printf("[400] [%s] %q bad request\n", r.RemoteAddr, r.URL.Path)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			log.Default().Printf("[500] [%s] %q error: %v\n", r.RemoteAddr, r.URL.Path, err)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	log.Default().Printf("[201] [%s] %q fruit set\n", r.RemoteAddr, r.URL.Path)
}

func user(w http.ResponseWriter, r *http.Request) {
	// handles URL: /fruit/:id - gets single user's fruit
	if corsPreflight(w, r) {
		return
	}

	if !validateMethod(w, r, "GET") {
		return
	}

	getByUser(w, r)

}

func root(w http.ResponseWriter, r *http.Request) {
	// handles URL: / - returns all fruits
	if corsPreflight(w, r) {
		return
	}

	if !validateMethod(w, r, "GET") {
		return
	}

	getAllFruits(r, w)

}
