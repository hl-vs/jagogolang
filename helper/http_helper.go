package helper

import (
	"encoding/json"
	"net/http"
)

func SetJSONError(v any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "error",
		"error":  s,
	})
}

func PrintJSONSuccess(v any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
