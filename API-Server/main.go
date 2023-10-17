package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct { //created a struct to store our values
	Id       int
	Name     string
	Quantity int
	Price    float64
}

var Products []Product //declared a variable of type struct called Product.

func main() {
	Products = []Product{
		Product{Id: 1, Name: "table", Quantity: 100, Price: 600.00},
		Product{Id: 2, Name: "chair", Quantity: 200, Price: 300.00},
	}
	handledRequests()

}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a web server API test")
	log.Println("Endpoint hit")
}
func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: returnAllProducts")
	json.NewEncoder(w).Encode(Products)
}
func handledRequests() {
	http.HandleFunc("/products", returnAllProducts)
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":8081", nil)
}
