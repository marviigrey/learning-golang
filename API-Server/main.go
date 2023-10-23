package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Product struct { //created a struct to store our values
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product //declared a variable of type struct called Product.

func main() {

	Products = []Product{
		Product{Id: "1", Name: "table", Quantity: 100, Price: 600.00},
		Product{Id: "2", Name: "chair", Quantity: 200, Price: 300.00},
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

func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, product := range Products {
		if string(product.Id) == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func handledRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", returnAllProducts)
	myRouter.HandleFunc("/product/{id}", getProduct)
	myRouter.HandleFunc("/", homepage)
	http.ListenAndServe(":8083", myRouter)
}
