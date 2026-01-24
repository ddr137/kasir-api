package handlers

import (
	"encoding/json"
	"net/http"

	"kasir-api/models"
	"kasir-api/utils"
)

// GET /health
func Health(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, map[string]string{
		"status":  "OK",
		"message": "API Running",
	})
}

// GET /api/products
func ListProducts(w http.ResponseWriter, r *http.Request) {
	models.Mu.RLock()
	defer models.Mu.RUnlock()
	utils.RespondJSON(w, http.StatusOK, models.Products)
}

// POST /api/products
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if product.Name == "" {
		utils.RespondError(w, http.StatusBadRequest, "Name is required")
		return
	}
	if product.Price <= 0 {
		utils.RespondError(w, http.StatusBadRequest, "Price must be greater than 0")
		return
	}
	if product.Stock < 0 {
		utils.RespondError(w, http.StatusBadRequest, "Stock cannot be negative")
		return
	}

	models.Mu.Lock()
	product.ID = models.GetNextID()
	models.Products = append(models.Products, product)
	models.Mu.Unlock()

	utils.RespondJSON(w, http.StatusCreated, product)
}

// GET /api/products/{id}
func GetProduct(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDFromRequest(r, w)
	if !ok {
		return
	}

	models.Mu.RLock()
	found, _ := models.FindProductByID(id)
	models.Mu.RUnlock()

	if found == nil {
		utils.RespondError(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, found)
}

// PUT /api/products/{id}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDFromRequest(r, w)
	if !ok {
		return
	}

	models.Mu.Lock()
	defer models.Mu.Unlock()

	found, _ := models.FindProductByID(id)
	if found == nil {
		utils.RespondError(w, http.StatusNotFound, "Product not found")
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if product.Name == "" {
		utils.RespondError(w, http.StatusBadRequest, "Name is required")
		return
	}
	if product.Price <= 0 {
		utils.RespondError(w, http.StatusBadRequest, "Price must be greater than 0")
		return
	}
	if product.Stock < 0 {
		utils.RespondError(w, http.StatusBadRequest, "Stock cannot be negative")
		return
	}

	found.Name = product.Name
	found.Price = product.Price
	found.Stock = product.Stock

	utils.RespondJSON(w, http.StatusOK, found)
}

// DELETE /api/products/{id}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDFromRequest(r, w)
	if !ok {
		return
	}

	models.Mu.Lock()
	defer models.Mu.Unlock()

	_, foundIndex := models.FindProductByID(id)
	if foundIndex == -1 {
		utils.RespondError(w, http.StatusNotFound, "Product not found")
		return
	}

	models.Products = append(models.Products[:foundIndex], models.Products[foundIndex+1:]...)

	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"message":  "Product deleted successfully",
		"products": models.Products,
	})
}
