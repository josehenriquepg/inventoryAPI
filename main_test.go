package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	product := Product{Name: "Test Product", Description: "Test Description", Price: 10.0, Quantity: 1}
	jsonProduct, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonProduct))
	req.Header.Set("content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createProduct)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var responseProduct Product
	json.NewDecoder(rr.Body).Decode(&responseProduct)
	if responseProduct.Name != product.Name {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), jsonProduct)
	}
}

func TestGetProducts(t *testing.T) {
	req, _ := http.NewRequest("GET", "/products", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getProducts)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetProduct(t *testing.T) {
	product := Product{ID: 1, Name: "Test Product", Description: "Test Description", Price: 10.0, Quantity: 1}
	products = append(products, product)

	req, _ := http.NewRequest("GET", "/products/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getProduct)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestUpdateProduct(t *testing.T) {
	product := Product{ID: 1, Name: "Test Product", Description: "Test Description", Price: 10.0, Quantity: 1}
	products = append(products, product)

	updatedProduct := Product{Name: "Updated Product", Description: "Updated Description", Price: 20.0, Quantity: 2}
	jsonProduct, _ := json.Marshal(updatedProduct)
	req, _ := http.NewRequest("PUT", "/products/1", bytes.NewBuffer(jsonProduct))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateProduct)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteProduct(t *testing.T) {
	product := Product{ID: 1, Name: "Test Product", Description: "Test Description", Price: 10.0, Quantity: 1}
	products = append(products, product)

	req, _ := http.NewRequest("DELETE", "/products/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteProduct)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
