package main

var Route = struct {
	API struct {
		Product  string
		Category string
		Health   string
		APIDOC   string
	}
	ROOT string
}{
	API: struct {
		Product  string
		Category string
		Health   string
		APIDOC   string
	}{
		Product:  "/api/produk/",
		Category: "/api/categories/",
		Health:   "/health",
		APIDOC:   "/api-doc",
	},
	ROOT: "/",
}
