package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"kasir-api/models"
	"kasir-api/services"
	"kasir-api/utils"
)

type CategoryHandler struct {
	service services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

// ListCategories godoc
// @Summary      Show all categories
// @Description  Get all categories with pagination
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        page      query     int  false  "Page number" default(1)
// @Param        page_size query     int  false  "Page size" default(10)
// @Success      200       {object}  utils.APIResponse{data=[]models.Category}
// @Failure      500       {object}  utils.APIResponse
// @Router       /api/categories [get]
func (h *CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

	categories, meta, err := h.service.GetAllCategories(page, pageSize)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseSuccessWithMeta(w, "Categories retrieved successfully", categories, meta)
}

// CreateCategory godoc
// @Summary      Create a new category
// @Description  Create a new category with the input payload
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category  body      models.Category  true  "Category"
// @Success      201       {object}  utils.APIResponse{data=models.Category}
// @Failure      400       {object}  utils.APIResponse
// @Failure      500       {object}  utils.APIResponse
// @Router       /api/categories [post]
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if category.Name == "" {
		utils.ResponseError(w, http.StatusBadRequest, "Name is required")
		return
	}

	createdCategory, err := h.service.CreateCategory(category)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseCreated(w, "Category created successfully", createdCategory)
}

// GetCategory godoc
// @Summary      Get a category
// @Description  Get category by ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Success      200  {object}  utils.APIResponse{data=models.Category}
// @Failure      400  {object}  utils.APIResponse
// @Failure      404  {object}  utils.APIResponse
// @Failure      500  {object}  utils.APIResponse
// @Router       /api/categories/{id} [get]
func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDFromRequest(r, w)
	if !ok {
		return
	}

	category, err := h.service.GetCategoryByID(id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if category == nil {
		utils.ResponseError(w, http.StatusNotFound, "Category not found")
		return
	}

	utils.ResponseSuccess(w, "Category retrieved successfully", category)
}

// UpdateCategory godoc
// @Summary      Update a category
// @Description  Update category by ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id        path      int              true  "Category ID"
// @Param        category  body      models.Category  true  "Category"
// @Success      200       {object}  utils.APIResponse{data=models.Category}
// @Failure      400       {object}  utils.APIResponse
// @Failure      404       {object}  utils.APIResponse
// @Failure      500       {object}  utils.APIResponse
// @Router       /api/categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDFromRequest(r, w)
	if !ok {
		return
	}

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if category.Name == "" {
		utils.ResponseError(w, http.StatusBadRequest, "Name is required")
		return
	}

	updatedCategory, err := h.service.UpdateCategory(id, category)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if updatedCategory == nil {
		utils.ResponseError(w, http.StatusNotFound, "Category not found")
		return
	}

	utils.ResponseSuccess(w, "Category updated successfully", updatedCategory)
}

// DeleteCategory godoc
// @Summary      Delete a category
// @Description  Delete category by ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Success      200  {object}  utils.APIResponse
// @Failure      400  {object}  utils.APIResponse
// @Failure      500  {object}  utils.APIResponse
// @Router       /api/categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDFromRequest(r, w)
	if !ok {
		return
	}

	err := h.service.DeleteCategory(id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseSuccess(w, "Category deleted successfully", nil)
}
