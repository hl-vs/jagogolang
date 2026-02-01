package helper

var Route = struct {
	API struct {
		Product      string
		ProductByID  string
		Category     string
		CategoryByID string
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
		Health       string
		APIDOC       string
	}{
		Product:      "/api/product",
		ProductByID:  "/api/product/",
		Category:     "/api/category",
		CategoryByID: "/api/category/",
		Health:       "/health",
		APIDOC:       "/api-doc",
	},
	ROOT: "/",
}
