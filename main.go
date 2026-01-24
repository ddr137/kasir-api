package main

import (
	"fmt"
	"log"
	"net/http"

	"kasir-api/handlers"
)

func main() {
	http.HandleFunc("/health", handlers.Health)
	http.HandleFunc("GET /api/products", handlers.ListProducts)
	http.HandleFunc("POST /api/products", handlers.CreateProduct)
	http.HandleFunc("GET /api/products/{id}", handlers.GetProduct)
	http.HandleFunc("PUT /api/products/{id}", handlers.UpdateProduct)
	http.HandleFunc("DELETE /api/products/{id}", handlers.DeleteProduct)

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
