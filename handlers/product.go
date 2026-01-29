package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"kasir-api/models"
	"kasir-api/services"
	"kasir-api/utils"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

// GET /health
func Health(w http.ResponseWriter, r *http.Request) {
	utils.ResponseSuccess(w, "API Running", nil)
}

// ListProducts godoc
// @Summary      Show all products
// @Description  Get all products with pagination
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page      query     int  false  "Page number" default(1)
// @Param        page_size query     int  false  "Page size" default(10)
// @Success      200       {object}  utils.APIResponse{data=[]models.Product}
// @Failure      500       {object}  utils.APIResponse
// @Router       /api/products [get]
func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

	products, meta, err := h.service.GetAllProducts(page, pageSize)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseSuccessWithMeta(w, "Products retrieved successfully", products, meta)
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Create a new product with the input payload
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      models.Product  true  "Product"
// @Success      201      {object}  utils.APIResponse{data=models.Product}
// @Failure      400      {object}  utils.APIResponse
// @Failure      500      {object}  utils.APIResponse
// @Router       /api/products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if product.Name == "" {
		utils.ResponseError(w, http.StatusBadRequest, "Name is required")
		return
	}
	if product.Price <= 0 {
		utils.ResponseError(w, http.StatusBadRequest, "Price must be greater than 0")
		return
	}
	if product.Stock < 0 {
		utils.ResponseError(w, http.StatusBadRequest, "Stock cannot be negative")
		return
	}
	if product.CategoryID <= 0 {
		utils.ResponseError(w, http.StatusBadRequest, "Category ID is required")
		return
	}

	createdProduct, err := h.service.CreateProduct(product)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseCreated(w, "Product created successfully", createdProduct)
}

// GetProduct godoc
// @Summary      Get a product
// @Description  Get product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  utils.APIResponse{data=models.Product}
// @Failure      400  {object}  utils.APIResponse
// @Failure      404  {object}  utils.APIResponse
// @Failure      500  {object}  utils.APIResponse
// @Router       /api/products/{id} [get]
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDFromRequest(r, w)
	if !ok {
		return
	}

	product, err := h.service.GetProductByID(id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if product == nil {
		utils.ResponseError(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.ResponseSuccess(w, "Product retrieved successfully", product)
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Update product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      int             true  "Product ID"
// @Param        product  body      models.Product  true  "Product"
// @Success      200      {object}  utils.APIResponse{data=models.Product}
// @Failure      400      {object}  utils.APIResponse
// @Failure      404  {object}  utils.APIResponse
// @Failure      500  {object}  utils.APIResponse
// @Router       /api/products/{id} [put]
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDFromRequest(r, w)
	if !ok {
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Basic validation
	if product.Name == "" {
		utils.ResponseError(w, http.StatusBadRequest, "Name is required")
		return
	}
	if product.Price <= 0 {
		utils.ResponseError(w, http.StatusBadRequest, "Price must be greater than 0")
		return
	}
	if product.Stock < 0 {
		utils.ResponseError(w, http.StatusBadRequest, "Stock cannot be negative")
		return
	}
	if product.CategoryID <= 0 {
		utils.ResponseError(w, http.StatusBadRequest, "Category ID is required")
		return
	}

	updatedProduct, err := h.service.UpdateProduct(id, product)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if updatedProduct == nil {
		utils.ResponseError(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.ResponseSuccess(w, "Product updated successfully", updatedProduct)
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  utils.APIResponse
// @Failure      400  {object}  utils.APIResponse
// @Failure      500  {object}  utils.APIResponse
// @Router       /api/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDFromRequest(r, w)
	if !ok {
		return
	}

	err := h.service.DeleteProduct(id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseSuccess(w, "Product deleted successfully", nil)
}
