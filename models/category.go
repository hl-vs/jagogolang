package models

type Category struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
}

// var category = []Category{
// 	{ID: 1, Nama: "Makanan", Deskripsi: "Kategori Makanan"},
// 	{ID: 2, Nama: "Minuman", Deskripsi: "Kategori Minuman"},
// }

// // GET localhost:8080/api/categories/
// // GET localhost:8080/api/categories/{id}
// func listCategory(w http.ResponseWriter, r *http.Request) {
// 	idStr := strings.TrimPrefix(r.URL.Path, Route.API.Category)
// 	if idStr == "" {
// 		printJSONSuccess(category, w)
// 	} else {
// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			setJSONError(fmt.Sprintf("Invalid Category ID: %s", idStr), w)
// 			return
// 		}

// 		// cari category berdasarkan id
// 		for _, p := range category {
// 			if p.ID == id {
// 				printJSONSuccess(p, w)
// 				return
// 			}
// 		}

// 		setJSONNotFound(fmt.Sprintf("Tidak ditemukan Category ID: %d", id), w)
// 		return
// 	}
// }

// func addCategory(w http.ResponseWriter, r *http.Request) {
// 	// get data dari request
// 	var newCategory Category
// 	err := json.NewDecoder(r.Body).Decode(&newCategory)
// 	if err != nil {
// 		setJSONError("Category tidak valid.", w)
// 		return
// 	}

// 	// validasi data
// 	if newCategory.Nama == "" || newCategory.Deskripsi == "" {
// 		setJSONError("Nama, Deskripsi wajib diisi dengan benar.", w)
// 		return
// 	}

// 	// cek nama category sudah ada atau belum
// 	for _, p := range category {
// 		if strings.EqualFold(p.Nama, newCategory.Nama) {
// 			setJSONError(fmt.Sprintf("Nama category '%s' sudah ada.", newCategory.Nama), w)
// 			return
// 		}
// 	}

// 	// set id baru
// 	newID := 1
// 	if len(category) > 0 {
// 		newID = category[len(category)-1].ID + 1
// 	}
// 	newCategory.ID = newID

// 	// tambahkan ke slice
// 	category = append(category, newCategory)

// 	// kembalikan response
// 	printJSONSuccess(map[string]any{
// 		"status":  "success",
// 		"message": fmt.Sprintf("'%s' (ID: %d) berhasil disimpan.", newCategory.Nama, newID),
// 		"data":    newCategory,
// 	}, w)
// }

// func updateCategory(w http.ResponseWriter, r *http.Request) {
// 	idStr := strings.TrimPrefix(r.URL.Path, Route.API.Category)
// 	if idStr == "" {
// 		setJSONNotFound("ID Category tidak ditemukan", w)
// 		return
// 	} else {
// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			setJSONError("ID Category tidak valid", w)
// 			return
// 		}

// 		// validasi data
// 		var updateCategory Category
// 		parseErr := json.NewDecoder(r.Body).Decode(&updateCategory)

// 		if parseErr != nil {
// 			setJSONError("Data category tidak valid", w)
// 			return
// 		}

// 		for i := range category {
// 			if category[i].ID == id {

// 				updateCategory.ID = id
// 				category[i] = updateCategory

// 				printJSONSuccess(map[string]any{
// 					"status":  "success",
// 					"message": fmt.Sprintf("Data category ID:%d telah di-update", updateCategory.ID),
// 					"data":    category[i],
// 				}, w)
// 				return
// 			}
// 		}
// 		setJSONNotFound("ID Category tidak ditemukan", w)
// 	}
// }

// func deleteCategory(w http.ResponseWriter, r *http.Request) {
// 	idStr := strings.TrimPrefix(r.URL.Path, Route.API.Category)
// 	if idStr == "" {
// 		setJSONNotFound("ID Category tidak ditemukan", w)
// 		return
// 	} else {
// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			setJSONError("ID Category tidak valid", w)
// 			return
// 		}

// 		for i, p := range category {
// 			if category[i].ID == id {
// 				category = append(category[:i], category[i+1:]...)

// 				printJSONSuccess(map[string]any{
// 					"status":  "success",
// 					"message": fmt.Sprintf("Category '%s' berhasil dihapus", p.Nama),
// 				}, w)
// 				return
// 			}
// 		}

// 		setJSONNotFound(fmt.Sprintf("Category ID '%d' tidak ditemukan", id), w)
// 	}
// }
