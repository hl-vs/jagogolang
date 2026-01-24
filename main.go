package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var Route = struct {
	API struct {
		Product string
		Health  string
	}
	ROOT string
}{
	API: struct {
		Product string
		Health  string
	}{
		Product: "/api/produk/",
		Health:  "/health",
	},
	ROOT: "/",
}

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

var produk = []Produk{
	{ID: 1, Nama: "Indomie", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
}

func setJSONError(s string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{
		"error": s,
	})
}

func setJSONNotFound(s string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": s,
	})
}

func printJSONSuccess(v any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// GET localhost:8080/api/produk/
// GET localhost:8080/api/produk/{id}
func listProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, Route.API.Product)
	if idStr == "" {
		printJSONSuccess(produk, w)
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			setJSONError(fmt.Sprintf("Invalid Produk ID: %s", idStr), w)
			return
		}

		for _, p := range produk {
			if p.ID == id {
				printJSONSuccess(p, w)
				return
			}
		}

		setJSONNotFound(fmt.Sprintf("Tidak ditemukan Produk ID: %d", id), w)
		return
	}
}

func main() {
	http.HandleFunc(Route.API.Product, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			listProduk(w, r)
		}

		if r.Method == "POST" {

		}

		if r.Method == "UPDATE" {

		}

		if r.Method == "PATCH" {

		}

		if r.Method == "PUT" {

		}
	})

	http.HandleFunc(Route.API.Health, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "API Running",
		})
	})

	http.HandleFunc(Route.ROOT, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Selamat Datang di Toko Online - API v1.0",
		})
	})

	fmt.Println("=======================")
	fmt.Println("Toko Online - API v1.0")
	fmt.Println("=======================")
	fmt.Println("API url: http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
