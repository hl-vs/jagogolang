package helper

var Route = struct {
	API struct {
		Product      string
		ProductByID  string
		Category     string
		CategoryByID string
		Checkout     string
		ReportToday  string
		ReportRange  string
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
		ReportToday  string
		ReportRange  string
		Health       string
		APIDOC       string
	}{
		Product:      "/api/product",
		ProductByID:  "/api/product/",
		Category:     "/api/category",
		CategoryByID: "/api/category/",
		Checkout:     "/api/checkout",
		ReportToday:  "/api/report/hari-ini",
		ReportRange:  "/api/report",
		Health:       "/health",
		APIDOC:       "/api-doc",
	},
	ROOT: "/",
}
