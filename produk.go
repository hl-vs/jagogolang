package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

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

		// cari produk berdasarkan id
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

func addProduk(w http.ResponseWriter, r *http.Request) {
	// get data dari request
	var newProduk Produk
	err := json.NewDecoder(r.Body).Decode(&newProduk)
	if err != nil {
		setJSONError("Produk tidak valid.", w)
		return
	}

	// validasi data
	if newProduk.Nama == "" || newProduk.Harga <= 0 || newProduk.Stok < 0 {
		setJSONError("Nama, Harga, dan Stok wajib diisi dengan benar.", w)
		return
	}

	// cek nama produk sudah ada atau belum
	for _, p := range produk {
		if strings.EqualFold(p.Nama, newProduk.Nama) {
			setJSONError(fmt.Sprintf("Produk dengan nama '%s' sudah ada.", newProduk.Nama), w)
			return
		}
	}

	// set id baru
	newID := 1
	if len(produk) > 0 {
		newID = produk[len(produk)-1].ID + 1
	}
	newProduk.ID = newID

	// tambahkan ke slice
	produk = append(produk, newProduk)

	// kembalikan response
	printJSONSuccess(map[string]any{
		"status":  "success",
		"message": fmt.Sprintf("'%s' (ID: %d) berhasil disimpan.", newProduk.Nama, newID),
		"data":    newProduk,
	}, w)
}

func updateProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, Route.API.Product)
	if idStr == "" {
		setJSONNotFound("Produk ID tidak ditemukan", w)
		return
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			setJSONError("Produk ID tidak valid", w)
			return
		}
		for i := range produk {
			if produk[i].ID == id {
				var updateProduk Produk
				parseErr := json.NewDecoder(r.Body).Decode(&updateProduk)

				if parseErr != nil {
					setJSONError("Produk tidak valid", w)
				}

				updateProduk.ID = id
				produk[i] = updateProduk

				printJSONSuccess(map[string]any{
					"status":  "success",
					"message": fmt.Sprintf("Data produk ID:%d telah di-update", updateProduk.ID),
					"data":    produk[i],
				}, w)
				return
			}
		}
		setJSONNotFound("Produk ID tidak ditemukan", w)
	}
}

func deleteProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, Route.API.Product)
	if idStr == "" {
		setJSONNotFound("Produk ID tidak ditemukan", w)
		return
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			setJSONError("Produk ID tidak valid", w)
			return
		}

		for i, p := range produk {
			if produk[i].ID == id {
				produk = append(produk[:i], produk[i+1:]...)

				printJSONSuccess(map[string]any{
					"status":  "success",
					"message": fmt.Sprintf("%s berhasil dihapus", p.Nama),
				}, w)
				return
			}
		}

		setJSONNotFound("Produk ID tidak ditemukan", w)
	}
}
