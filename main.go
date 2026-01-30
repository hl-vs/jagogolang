package main

import (
	"fmt"
	"kasir-online/handlers"
	"kasir-online/helper"
	"kasir-online/repositories"
	"kasir-online/services"
	"log"
	"net/http"
)

func main() {
	productRepo := repositories.NewProductRepository(nil)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	http.HandleFunc(helper.Route.API.ProductByID, productHandler.HandleByID)
	http.HandleFunc(helper.Route.API.Product, productHandler.HandleProducts)
	// http.HandleFunc(helper.Route.API.Product, func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "GET" {
	// 		models.listProduk(w, r)
	// 	}

	// 	if r.Method == "POST" {
	// 		addProduk(w, r)
	// 	}

	// 	if r.Method == "PUT" {
	// 		updateProduk(w, r)
	// 	}

	// 	if r.Method == "DELETE" {
	// 		deleteProduk(w, r)
	// 	}
	// })

	// http.HandleFunc(helper.Route.API.Category, func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "GET" {
	// 		listCategory(w, r)
	// 	}

	// 	if r.Method == "POST" {
	// 		addCategory(w, r)
	// 	}

	// 	if r.Method == "PUT" {
	// 		updateCategory(w, r)
	// 	}

	// 	if r.Method == "DELETE" {
	// 		deleteCategory(w, r)
	// 	}
	// })

	http.HandleFunc(helper.Route.API.Health, func(w http.ResponseWriter, r *http.Request) {
		helper.PrintJSONSuccess(map[string]string{
			"status":  "ok",
			"message": "API Running",
		}, w)
	})

	http.HandleFunc(helper.Route.API.APIDOC, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api-doc/JagoGolang/index.html")
	})

	http.HandleFunc(helper.Route.ROOT, func(w http.ResponseWriter, r *http.Request) {
		helper.PrintJSONSuccess(map[string]string{
			"message": "Selamat Datang di Kasir Online - API v1.0",
			"debug":   fmt.Sprintf("path=%s", r.URL.Path),
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
