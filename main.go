package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func setJSONError(v any, w http.ResponseWriter) {
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

func setJSONNotFound(s string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "error",
		"error":  s,
	})
}

func printJSONSuccess(v any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func main() {
	http.HandleFunc(Route.API.Product, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			listProduk(w, r)
		}

		if r.Method == "POST" {
			addProduk(w, r)
		}

		if r.Method == "PUT" {
			updateProduk(w, r)
		}

		if r.Method == "DELETE" {
			deleteProduk(w, r)
		}
	})

	http.HandleFunc(Route.API.Category, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			listCategory(w, r)
		}

		if r.Method == "POST" {
			addCategory(w, r)
		}

		if r.Method == "PUT" {
			updateCategory(w, r)
		}

		if r.Method == "DELETE" {
			deleteCategory(w, r)
		}
	})

	http.HandleFunc(Route.API.Health, func(w http.ResponseWriter, r *http.Request) {
		printJSONSuccess(map[string]string{
			"status":  "ok",
			"message": "API Running",
		}, w)
	})

	http.HandleFunc(Route.API.APIDOC, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api-doc/JagoGolang/index.html")
	})

	http.HandleFunc(Route.ROOT, func(w http.ResponseWriter, r *http.Request) {
		printJSONSuccess(map[string]string{
			"message": "Selamat Datang di Kasir Online - API v1.0",
		}, w)
	})

	fmt.Println("=======================")
	fmt.Println("Kasir Online - API v1.0")
	fmt.Println("=======================")
	fmt.Println("API url: http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
