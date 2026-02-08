package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

type WelcomeResponse struct {
	Message string `json:"message"`
	APIDoc  string `json:"api-doc"`
	Request string `json:"request"`
}

func SetJSONError(v any, w http.ResponseWriter) {
	status := http.StatusBadRequest
	log.Printf("--- Response: %d", status)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if s, ok := v.(string); ok {
		json.NewEncoder(w).Encode(map[string]string{
			"status": "error",
			"error":  s,
		})
		return
	}
	json.NewEncoder(w).Encode(v)
}

func SetJSONNotFound(s string, w http.ResponseWriter) {
	status := http.StatusNotFound
	log.Printf("--- Response: %d", status)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "error",
		"error":  s,
	})
}

func PrintJSONSuccess(v any, w http.ResponseWriter) {
	status := http.StatusOK
	log.Printf("--- Response: %d", status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
