package handlers

import (
	"encoding/json"
	"fmt"
	"kasir-online/helper"
	"kasir-online/models"
	"kasir-online/services"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CategoryHandler) HandleByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.PrintJSONSuccess(map[string]any{
		"total": len(categories),
		"data":  categories,
	}, w)
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, helper.Route.API.CategoryByID)
	if idStr == "" {
		h.GetAll(w, r)
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			helper.SetJSONError(fmt.Sprintf("Invalid Category ID: %s", idStr), w)
			return
		}
		category, err := h.service.GetByID(id)
		if err != nil {
			helper.SetJSONNotFound(fmt.Sprintf("Tidak ditemukan Category ID: %d", id), w)
			return
		}

		helper.PrintJSONSuccess(category, w)
		return
	}
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, helper.Route.API.CategoryByID)
	if idStr == "" {
		helper.SetJSONNotFound("Category ID tidak ditemukan", w)
		return
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			helper.SetJSONError("Category ID tidak valid", w)
			return
		}

		// validasi data
		var updateCategory models.Category
		err = json.NewDecoder(r.Body).Decode(&updateCategory)

		if err != nil {
			helper.SetJSONError(err.Error(), w)
			return
		}

		err = h.service.Update(id, &updateCategory)
		if err != nil {
			helper.SetJSONError(err.Error(), w)
			return
		}
		updateCategory.ID = id

		helper.PrintJSONSuccess(map[string]any{
			"status":  "success",
			"message": fmt.Sprintf("Data Category ID:%d telah di-update", updateCategory.ID),
			"data":    updateCategory,
		}, w)
		return
	}
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, helper.Route.API.CategoryByID)
	if idStr == "" {
		helper.SetJSONNotFound("Category ID tidak ditemukan", w)
		return
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			helper.SetJSONError("Category ID tidak valid", w)
			return
		}
		err = h.service.Delete(id)
		if err != nil {
			helper.SetJSONError(fmt.Sprintf("Category ID:%d tidak ditemukan", id), w)
		} else {
			helper.PrintJSONSuccess(map[string]any{
				"status":  "success",
				"message": fmt.Sprintf("Category ID:%d berhasil dihapus", id),
			}, w)
		}
		return
	}
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		helper.SetJSONError("Category tidak valid.", w)
		return
	}

	// validasi data
	if newCategory.Name == "" || newCategory.Description == "" {
		helper.SetJSONError("Name, dan Description wajib diisi dengan benar.", w)
		return
	}

	// tambahkan ke slice
	newID, err := h.service.Create(&newCategory)
	if err != nil {
		helper.SetJSONError(err.Error(), w)
		return
	}

	if newID != nil {
		newCategory.ID = *newID
	}

	// kembalikan response
	helper.PrintJSONSuccess(map[string]any{
		"status":  "success",
		"message": fmt.Sprintf("'%s' (ID: %d) berhasil disimpan.", newCategory.Name, *newID),
		"data":    newCategory,
	}, w)

}
