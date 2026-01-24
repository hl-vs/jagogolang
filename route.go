package main

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
