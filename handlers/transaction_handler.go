package handlers

import (
	"encoding/json"
	"kasir-online/helper"
	"kasir-online/models"
	"kasir-online/services"
	"net/http"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helper.SetJSONError(map[string]string{
			"status": "error",
			"error":  err.Error(),
		}, w)
		return
	}

	// validasi data
	for _, it := range req.Items {
		if it.ProductID <= 0 {
			helper.SetJSONError("Invalid ProductID", w)
			return
		}
		if it.Quantity <= 0 {
			helper.SetJSONError("Invalid Quantity", w)
			return
		}
	}

	transaction, err := h.service.Checkout(req.Items, true)
	if err != nil {
		helper.SetJSONError(map[string]string{
			"status": "error",
			"error":  err.Error(),
		}, w)
		return
	}

	helper.PrintJSONSuccess(transaction, w)
}
