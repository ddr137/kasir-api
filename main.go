package main

import (
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"kasir-api/utils"
	"log"
	"net/http"

	_ "kasir-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Kasir API
// @version         1.0
// @description     A simple POS (Point of Sale) API.
// @host            localhost:8093
// @BasePath        /

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := database.InitDB(config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	defer db.Close()

	// Dependency Injection - Product
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Dependency Injection - Category
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	http.HandleFunc("/health", handlers.Health)

	// Swagger
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Product Routes
	http.HandleFunc("GET /api/products", productHandler.ListProducts)
	http.HandleFunc("POST /api/products", productHandler.CreateProduct)
	http.HandleFunc("GET /api/products/{id}", productHandler.GetProduct)
	http.HandleFunc("PUT /api/products/{id}", productHandler.UpdateProduct)
	http.HandleFunc("DELETE /api/products/{id}", productHandler.DeleteProduct)

	// Category Routes
	http.HandleFunc("GET /api/categories", categoryHandler.ListCategories)
	http.HandleFunc("POST /api/categories", categoryHandler.CreateCategory)
	http.HandleFunc("GET /api/categories/{id}", categoryHandler.GetCategory)
	http.HandleFunc("PUT /api/categories/{id}", categoryHandler.UpdateCategory)
	http.HandleFunc("DELETE /api/categories/{id}", categoryHandler.DeleteCategory)

	fmt.Printf("Server running on http://localhost:%s\n", config.Port)
	if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
		log.Fatal(err)
	}
}
