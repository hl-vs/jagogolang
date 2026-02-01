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

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) HandleByID(w http.ResponseWriter, r *http.Request) {
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

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.PrintJSONSuccess(map[string]any{
		"total": len(products),
		"data":  products,
	}, w)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, helper.Route.API.ProductByID)
	if idStr == "" {
		h.GetAll(w, r)
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			helper.SetJSONError(fmt.Sprintf("Invalid Product ID: %s", idStr), w)
			return
		}
		product, err := h.service.GetByID(id)
		if err != nil {
			helper.SetJSONNotFound(fmt.Sprintf("Tidak ditemukan Product ID: %d", id), w)
			return
		}

		helper.PrintJSONSuccess(product, w)
		return
	}
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, helper.Route.API.ProductByID)
	if idStr == "" {
		helper.SetJSONNotFound("Product ID tidak ditemukan", w)
		return
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			helper.SetJSONError("Product ID tidak valid", w)
			return
		}

		// validasi data
		var updateProduct models.Product
		parseErr := json.NewDecoder(r.Body).Decode(&updateProduct)

		if parseErr != nil {
			helper.SetJSONError("Product tidak valid", w)
			return
		}

		err = h.service.Update(id, &updateProduct)
		if err != nil {
			helper.SetJSONError(err, w)
			return
		}
		updateProduct.ID = id

		helper.PrintJSONSuccess(map[string]any{
			"status":  "success",
			"message": fmt.Sprintf("Data Product ID:%d telah di-update", updateProduct.ID),
			"data":    updateProduct,
		}, w)
		return
	}
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, helper.Route.API.ProductByID)
	if idStr == "" {
		helper.SetJSONNotFound("Product ID tidak ditemukan", w)
		return
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			helper.SetJSONError("Product ID tidak valid", w)
			return
		}
		err = h.service.Delete(id)
		if err != nil {
			helper.SetJSONError(fmt.Sprintf("Product ID:%d tidak ditemukan", id), w)
		} else {
			helper.PrintJSONSuccess(map[string]any{
				"status":  "success",
				"message": fmt.Sprintf("Product ID:%d berhasil dihapus", id),
			}, w)
		}
		return
	}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		helper.SetJSONError("Product tidak valid.", w)
		return
	}

	// validasi data
	if newProduct.Name == "" || newProduct.Price <= 0 || newProduct.Stock < 0 {
		helper.SetJSONError("Name, Price, dan Stock wajib diisi dengan benar.", w)
		return
	}

	// tambahkan ke slice
	newID, err := h.service.Create(&newProduct)
	if err != nil {
		helper.SetJSONError(err, w)
		return
	}

	if newID != nil {
		newProduct.ID = *newID
	}

	// kembalikan response
	helper.PrintJSONSuccess(map[string]any{
		"status":  "success",
		"message": fmt.Sprintf("'%s' (ID: %d) berhasil disimpan.", newProduct.Name, *newID),
		"data":    newProduct,
	}, w)

}

// =============== v1 ================
func (h *ProductHandler) GetAll_V1(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll_V1()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.PrintJSONSuccess(map[string]any{
		"total": len(products),
		"data":  products,
	}, w)
}

func (h *ProductHandler) GetByID_V1(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, helper.Route.API.ProductByID)
	if idStr == "" {
		h.GetAll_V1(w, r)
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			helper.SetJSONError(fmt.Sprintf("Invalid Product ID: %s", idStr), w)
			return
		}

		// cari product berdasarkan id
		products, err := h.service.GetAll_V1()
		for _, p := range products {
			if p.ID == id {
				helper.PrintJSONSuccess(p, w)
				return
			}
		}

		helper.SetJSONNotFound(fmt.Sprintf("Tidak ditemukan Product ID: %d", id), w)
		return
	}
}

func (h *ProductHandler) Update_V1(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, helper.Route.API.ProductByID)
	if idStr == "" {
		helper.SetJSONNotFound("Product ID tidak ditemukan", w)
		return
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			helper.SetJSONError("Product ID tidak valid", w)
			return
		}

		// validasi data
		var updateProduct models.Product
		parseErr := json.NewDecoder(r.Body).Decode(&updateProduct)

		if parseErr != nil {
			helper.SetJSONError("Product tidak valid", w)
			return
		}

		products, err := h.service.GetAll_V1()
		for i := range products {
			if products[i].ID == id {
				updateProduct.ID = id
				products[i] = updateProduct

				helper.PrintJSONSuccess(map[string]any{
					"status":  "success",
					"message": fmt.Sprintf("Data Product ID:%d telah di-update", updateProduct.ID),
					"data":    products[i],
				}, w)
				return
			}
		}
		helper.SetJSONNotFound("Product ID tidak ditemukan", w)
	}
}

func (h *ProductHandler) Delete_V1(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, helper.Route.API.ProductByID)
	if idStr == "" {
		helper.SetJSONNotFound("Product ID tidak ditemukan", w)
		return
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			helper.SetJSONError("Product ID tidak valid", w)
			return
		}
		err2 := h.service.Delete_V1(id)
		if err2 != nil {
			helper.SetJSONError(fmt.Sprintf("Product ID:%d tidak ditemukan", id), w)
		} else {
			helper.PrintJSONSuccess(map[string]any{
				"status":  "success",
				"message": fmt.Sprintf("Product ID:%d berhasil dihapus", id),
			}, w)
		}
		return
	}
}

func (h *ProductHandler) Create_V1(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		helper.SetJSONError("Product tidak valid.", w)
		return
	}

	// validasi data
	if newProduct.Name == "" || newProduct.Price <= 0 || newProduct.Stock < 0 {
		helper.SetJSONError("Name, Price, dan Stock wajib diisi dengan benar.", w)
		return
	}

	// cek nama product sudah ada atau belum
	products, err := h.service.GetAll_V1()
	for _, p := range products {
		if strings.EqualFold(p.Name, newProduct.Name) {
			helper.SetJSONError(fmt.Sprintf("Name product '%s' sudah ada.", newProduct.Name), w)
			return
		}
	}

	// set id baru
	newID := 1
	if len(products) > 0 {
		newID = products[len(products)-1].ID + 1
	}
	newProduct.ID = newID

	// tambahkan ke slice
	h.service.Create_V1(&newProduct)

	// kembalikan response
	helper.PrintJSONSuccess(map[string]any{
		"status":  "success",
		"message": fmt.Sprintf("'%s' (ID: %d) berhasil disimpan.", newProduct.Name, newID),
		"data":    newProduct,
	}, w)
}
