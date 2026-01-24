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
Server default mendengarkan di `http://localhost:8080`.

## Endpoint utama
- GET  /api/produk/        — daftar semua produk
- GET  /api/produk/{id}    — ambil produk berdasarkan ID
- POST /api/produk         — tambah produk baru
- PUT  /api/produk/{id}    — perbarui produk
- DELETE /api/produk/{id}  — hapus produk
- GET  /health             — health check
- GET  /                   — pesan selamat datang

## Contoh curl
- List:
```sh
curl http://localhost:8080/api/produk/
```
- Get by ID:
```sh
curl http://localhost:8080/api/produk/1
```
- Create:
```sh
curl -X POST -H "Content-Type: application/json" -d '{"nama":"Susu","harga":5000,"stok":10}' http://localhost:8080/api/produk
```
- Update:
```sh
curl -X PUT -H "Content-Type: application/json" -d '{"nama":"Susu UHT","harga":5500,"stok":8}' http://localhost:8080/api/produk/1
```
- Delete:
```sh
curl -X DELETE http://localhost:8080/api/produk/1
```

## Struktur file (ringkas)
- [go.mod](go.mod)
- [main.go](main.go)
- [produk.go](produk.go)
- [route.go](route.go)
- [settings.json](settings.json)
- [README.md](README.md)
- [.gitignore](.gitignore)
- [api-doc](api-doc/) — dokumentasi collection / YAML untuk pengujian API

## File penting
- [main.go](main.go) — entrypoint dan pendaftaran route
- [produk.go](produk.go) — model dan handler terkait produk
- [route.go](route.go) — konfigurasi route

Kontribusi diterima. Untuk pertanyaan atau instruksi menjalankan lebih lanjut, beri tahu saya.