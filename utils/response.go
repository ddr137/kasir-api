package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// PaginationMeta can be used for pagination details later
type PaginationMeta struct {
	Page      int `json:"page"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

type APIResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"` // Use interface{} to allow null
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ResponseSuccessWithMeta(w http.ResponseWriter, message string, data interface{}, meta interface{}) {
	response := APIResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    meta,
	}
	RespondJSON(w, http.StatusOK, response)
}

func ResponseSuccess(w http.ResponseWriter, message string, data interface{}) {
	ResponseSuccessWithMeta(w, message, data, nil)
}

func ResponseCreated(w http.ResponseWriter, message string, data interface{}) {
	response := APIResponse{
		Code:    http.StatusCreated,
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    nil,
	}
	RespondJSON(w, http.StatusCreated, response)
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	response := APIResponse{
		Code:    code,
		Status:  "error",
		Message: message,
		Data:    nil,
		Meta:    nil,
	}
	RespondJSON(w, code, response)
}

func ParseIDFromRequest(r *http.Request, w http.ResponseWriter) (int, bool) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "Invalid ID")
		return 0, false
	}
	return id, true
}
