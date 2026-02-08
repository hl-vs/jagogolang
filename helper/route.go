package helper

var Route = struct {
	API struct {
		Product      string
		ProductByID  string
		Category     string
		CategoryByID string
		Checkout     string
		Health       string
		APIDOC       string
	}
	ROOT string
}{
	API: struct {
		Product      string
		ProductByID  string
		Category     string
		CategoryByID string
		Checkout     string
		Health       string
		APIDOC       string
	}{
		Product:      "/api/product",
		ProductByID:  "/api/product/",
		Category:     "/api/category",
		CategoryByID: "/api/category/",
		Checkout:     "/api/checkout",
		Health:       "/health",
		APIDOC:       "/api-doc",
	},
	ROOT: "/",
}
