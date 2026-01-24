# Kasir Online API v1.0

Backend project untuk latihan di **JagoGolang Bootcamp**.

## Ringkasan

API sederhana untuk mengelola produk (CRUD) dengan endpoint HTTP.

## Persyaratan

- Go 1.18+

## Menjalankan

Jalankan server (direktori project):

```sh
go run .
```

Atau buat binary dan jalankan:

```sh
go build -o tmp/main .
./tmp/main
```

Server default berjalan di `http://localhost:8080`.

## Endpoint utama

### Produk

Model `Produk` memiliki field: `id`, `nama`, `harga`, `stok`
Endpoint:

- GET  /api/produk/        — daftar semua produk
- GET  /api/produk/{id}    — ambil produk berdasarkan ID
- POST /api/produk         — tambah produk baru
- PUT  /api/produk/{id}    — perbarui produk
- DELETE /api/produk/{id}  — hapus produk
- GET  /health             — health check
- GET  /                   — pesan selamat datang

### Category

Model `Category` memiliki field: `id`, `nama`, `deskripsi`.

Endpoint:

- GET  /api/categories/        — daftar semua kategori
- GET  /api/categories/{id}    — ambil kategori berdasarkan ID
- POST /api/categories         — tambah kategori baru
- PUT  /api/categories/{id}    — perbarui kategori
- DELETE /api/categories/{id}  — hapus kategori

## Contoh curl

### Curl Produk

- List produk:

```sh
curl http://localhost:8080/api/produk/
```

- Get produk by ID:

```sh
curl http://localhost:8080/api/produk/1
```

- Create produk:

```sh
curl -X POST -H "Content-Type: application/json" -d '{"nama":"Susu","harga":5000,"stok":10}' http://localhost:8080/api/produk
```

- Update produk:

```sh
curl -X PUT -H "Content-Type: application/json" -d '{"nama":"Susu UHT","harga":5500,"stok":8}' http://localhost:8080/api/produk/1
```

- Delete produk:

```sh
curl -X DELETE http://localhost:8080/api/produk/1
```

### Curl Category

- List categories:

```sh
curl http://localhost:8080/api/categories/
```

- Get category by ID:

```sh
curl http://localhost:8080/api/categories/1
```

- Create category:

```sh
curl -X POST -H "Content-Type: application/json" -d '{"nama":"Makanan","deskripsi":"Kategori Makanan"}' http://localhost:8080/api/categories
```

- Update category:

```sh
curl -X PUT -H "Content-Type: application/json" -d '{"nama":"Minuman","deskripsi":"Kategori Minuman"}' http://localhost:8080/api/categories/1
```

- Delete category:

```sh
curl -X DELETE http://localhost:8080/api/categories/1
```

## Struktur file

- [go.mod](go.mod)
- [settings.json](settings.json)
- [produk.go](produk.go)
- [main.go](main.go)
- [kategori.go](kategori.go)
- [route.go](route.go)
- [README.md](README.md)
- [.gitignore](.gitignore)
- [api-doc](api-doc/) — dokumentasi collection / YAML untuk pengujian API

## File penting

- [main.go](main.go) — entrypoint dan pendaftaran route
- [produk.go](produk.go) — model dan handler terkait produk
- [kategori.go](kategori.go) — model dan handler terkait kategori
- [route.go](route.go) — konfigurasi route

Untuk pertanyaan atau instruksi menjalankan lebih lanjut, beri tahu saya.
