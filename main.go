package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	// r.Handle("/status/", NotImplemented).Methods("GET")
	r.Handle("/status", NotImplemented).Methods("GET")

	http.ListenAndServe(":3000", r)

}

// NotImplemented is called for importing routes not implemented
var NotImplemented = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Still Alive!!"))
	})
