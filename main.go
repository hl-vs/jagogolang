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

	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// custom muxing for log
	mux := http.NewServeMux()

	// handler
	mux.HandleFunc(helper.Route.API.ProductByID, productHandler.HandleByID)
	mux.HandleFunc(helper.Route.API.Product, productHandler.HandleProducts)
	mux.HandleFunc(helper.Route.API.CategoryByID, categoryHandler.HandleByID)
	mux.HandleFunc(helper.Route.API.Category, categoryHandler.HandleCategories)
	mux.HandleFunc(helper.Route.API.Checkout, transactionHandler.HandleCheckout)
	mux.HandleFunc(helper.Route.API.ReportRange, transactionHandler.HandleReport)
	mux.HandleFunc(helper.Route.API.ReportToday, transactionHandler.HandleReport)

	mux.HandleFunc(helper.Route.API.Health, func(w http.ResponseWriter, r *http.Request) {
		helper.PrintJSONSuccess(map[string]string{
			"status":  "ok",
			"message": "API Running",
		}, w)
	})

	mux.HandleFunc(helper.Route.API.APIDOC, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api-doc/JagoGolang/index.html")
	})

	mux.HandleFunc(helper.Route.ROOT, func(w http.ResponseWriter, r *http.Request) {
		res := helper.WelcomeResponse{
			Message: "Selamat Datang di Kasir Online - API v1.0",
			APIDoc:  "/api-doc",
			Request: r.URL.Path,
		}
		helper.PrintJSONSuccess(res, w)
	})

	addr := ":" + config.Port

	fmt.Println("=======================")
	fmt.Println("Kasir Online - API v1.0")
	fmt.Println("=======================")
	fmt.Printf("API url: http://localhost:%s\n\n", config.Port)

	wrappedMux := loggingMiddleware(mux)

	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
