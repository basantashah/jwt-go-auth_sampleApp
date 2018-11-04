package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	// r.Handle("/status/", NotImplemented).Methods("GET")
	r.Handle("/status", NotImplemented).Methods("GET")
	r.Handle("/products", NotImplemented).Methods("GET")
	r.Handle("products{slug}/feedback", NotImplemented).Methods("POST")

	http.ListenAndServe(":3000", r)

}

// NotImplemented is called for importing routes not implemented
var NotImplemented = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Still Alive!!"))
	})

type Product struct {
	ID          int
	Name        string
	Slug        string
	Description string
}

var products = []Product{
	Product{ID: 1, Name: "cheese cake", Slug: "cheese cake", Description: "this is kind of cake made of cheese"},
	Product{ID: 8, Name: "pan cake", Slug: "cheese cake", Description: "this is kind of cake made of cheese"},
	Product{ID: 2, Name: "slum dog", Slug: "cheese cake", Description: "this is kind of cake made of cheese"},
	Product{ID: 3, Name: "nintendo switch", Slug: "cheese cake", Description: "this is kind of cake made of cheese"},
	Product{ID: 4, Name: "Paneer masala", Slug: "Paneer masala", Description: "this is kind of cake made of cheese"},
	Product{ID: 5, Name: "Namaste masala", Slug: "Namaste masala", Description: "this is kind of cake made of cheese"},
	Product{ID: 6, Name: "Distance Love", Slug: "Distance Love", Description: "this is kind of cake made of cheese"},
	Product{ID: 7, Name: "Chola Bhatoori", Slug: "Chola Bhatoori", Description: "this is kind of cake made of cheese"},
}

var StatusHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is up and Running"))
	})

var ProductsHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		payload, _ := json.Marshal(products)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(payload))
	})
