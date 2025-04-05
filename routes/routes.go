package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouters(router *mux.Router) {
	router.HandleFunc("/", HomeHandler).Methods("GET")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
