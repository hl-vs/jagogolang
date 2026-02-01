package main

import (
	"fmt"
	"kasir-online/database"
	"kasir-online/handlers"
	"kasir-online/helper"
	"kasir-online/repositories"
	"kasir-online/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"APP_PORT"`
	DBConn string `mapstructure:"DATABASE_URL"`
}

func main() {
	// setup viper config
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("APP_PORT"),
		DBConn: viper.GetString("DATABASE_URL"),
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	http.HandleFunc(helper.Route.API.ProductByID, productHandler.HandleByID)
	http.HandleFunc(helper.Route.API.Product, productHandler.HandleProducts)
	http.HandleFunc(helper.Route.API.CategoryByID, categoryHandler.HandleByID)
	http.HandleFunc(helper.Route.API.Category, categoryHandler.HandleCategories)

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
			"path":    fmt.Sprintf("%s", r.URL.Path),
		}, w)
	})

	addr := "0.0.0.0:" + config.Port

	fmt.Println("=======================")
	fmt.Println("Kasir Online - API v1.0")
	fmt.Println("=======================")
	fmt.Print("API url: http://localhost:", config.Port)
	fmt.Println("\n---")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
