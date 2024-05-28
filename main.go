package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

var products []Product
var nextID int

func main() {
	r := mux.NewRouter()

	// CRUD endpoints
	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	fmt.Println("Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// list products
func getProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

// create product
func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = nextID
	nextID++
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

// get product by ID
func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, product := range products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	http.NotFound(w, r)
}

// update product
func updateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, product := range products {
		if product.ID == id {
			products = append(products[:index], products[index+1:]...)
			var updateProduct Product
			_ = json.NewDecoder(r.Body).Decode(&updateProduct)
			updateProduct.ID = id
			products = append(products, updateProduct)
			json.NewEncoder(w).Encode(updateProduct)
			return
		}
	}
	http.NotFound(w, r)
}

// delete product
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, product := range products {
		if product.ID == id {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}
