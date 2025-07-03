# Order Management API

Order Management API adalah sistem backend untuk manajemen produk, pemesanan, dan autentikasi berbasis Go (Golang) dengan arsitektur clean code.

## Fitur Utama
- **Autentikasi JWT** (register, login, role admin/customer)
- **CRUD Produk** (khusus admin)
- **Order Produk** (customer, stok otomatis berkurang)
- **Riwayat Pesanan Customer**
- **Validasi & Error Handling**
- **Swagger API Documentation**

## Admin Default
Saat aplikasi pertama kali dijalankan, jika belum ada user dengan role `admin`, maka akan otomatis dibuat akun admin default:
- **Email:** `admin@gmail.com`
- **Password:** `admin123`

> **Penting:** Segera ganti password admin default ini di production!

## Struktur Project (Clean Code)
```
.
├── config/         # Koneksi database
├── docs/           # File Swagger auto-generated
├── handler/        # HTTP handler (controller)
├── middleware/     # Middleware (JWT, dsb)
├── models/         # Model database
├── repository/     # Query database
├── routes/         # Routing utama
├── service/        # Business logic
├── utils/          # Helper (response, dsb)
├── main.go         # Entrypoint
├── go.mod
├── .env.example    # Contoh konfigurasi environment
└── README.md
```

## Setup & Menjalankan
1. **Clone repo & install dependency**
   ```sh
   git clone https://github.com/wahyuutomoputra/order-management
   cd order-management
   go mod tidy
   ```
2. **Siapkan file .env di root project**
   - Copy file contoh:
     ```sh
     cp .env.example .env
     # Edit .env sesuai konfigurasi MySQL Anda
     ```
   - Contoh isi `.env`:
     ```env
     DB_USER=root
     DB_PASS=yourpassword
     DB_HOST=127.0.0.1:3306
     DB_NAME=orderdb
     PORT=8080
     ```
   - **Catatan:**
     - File `.env` harus ada di root project. Semua konfigurasi database dan port aplikasi akan otomatis diambil dari file ini saat aplikasi dijalankan.
     - Untuk mengubah port aplikasi, cukup ubah nilai `PORT` di `.env` (misal `PORT=9000`).
3. **Generate Swagger docs**
   ```sh
   swag init
   ```
4. **Jalankan aplikasi**
   ```sh
   go run main.go
   ```
5. **Akses dokumentasi Swagger**
   - [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
     (ganti 8080 sesuai port yang Anda set di .env)

## Contoh Endpoint
- `POST /register` — register user baru
- `POST /login` — login, dapatkan JWT
- `GET /me` — info user login
- `GET /products` — list produk
- `POST /admin/products` — tambah produk (admin)
- `POST /orders` — buat order (customer)
- `GET /orders/history` — riwayat order customer

## Response Format Seragam
```json
{
  "success": true,
  "data": { ... },
  "message": "..."
}
```
atau
```json
{
  "success": false,
  "error": "..."
}
```

## Swagger
- Semua endpoint terdokumentasi otomatis di Swagger.
- Untuk update dokumentasi, jalankan `swag init` setelah mengubah anotasi handler.

---

**Kontribusi & pertanyaan silakan buka issue atau pull request.** 