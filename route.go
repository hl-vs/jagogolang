package main

var Route = struct {
	API struct {
		Product  string
		Category string
		Health   string
	}
	ROOT string
}{
	API: struct {
		Product  string
		Category string
		Health   string
	}{
		Product:  "/api/produk/",
		Category: "/api/categories/",
		Health:   "/health",
	},
	ROOT: "/",
}
