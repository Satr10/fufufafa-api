# Fufufafa API

Sebuah API sederhana yang dibangun dengan Go (menggunakan framework Fiber) untuk menyajikan kutipan "Fufufafa". API ini terhubung ke database PostgreSQL untuk mengambil data kutipan.

## Fitur

* Menyediakan endpoint RESTful untuk mengakses kutipan.
* Mengambil semua kutipan yang tersedia.
* Mengambil kutipan spesifik berdasarkan ID.
* Mengambil kutipan acak.
* Terintegrasi dengan Swagger untuk dokumentasi API interaktif.
* Konfigurasi melalui variabel lingkungan (`.env.local`).
* Middleware untuk logging dan penanganan halaman tidak ditemukan (404).
* Dukungan untuk deployment di Vercel.
* Caching pada endpoint `/api`.

## Teknologi yang Digunakan

* **Bahasa:** Go
* **Framework Web:** [Fiber](https://gofiber.io/)
* **ORM:** [GORM](https://gorm.io/)
* **Database:** PostgreSQL
* **Dokumentasi API:** [Swagger](https://swagger.io/) (via [swaggo/swag](https://github.com/swaggo/swag) & [gofiber/contrib/swagger](https://github.com/gofiber/contrib/swagger))
* **Manajemen Konfigurasi:** [godotenv](https://github.com/joho/godotenv)

## Prasyarat

* Go (versi 1.18 atau lebih baru direkomendasikan)
* Server PostgreSQL yang sedang berjalan
* Git

## Instalasi & Menjalankan Lokal

1.  **Clone repository:**
    ```bash
    git clone https://github.com/Satr10/fufufafa-api
    cd fufufafa-api
    ```

2.  **Buat file `.env.local`:**
    Buat file bernama `.env.local` di root direktori proyek dan isi dengan konfigurasi koneksi database PostgreSQL Anda. Contoh:
    ```env
    # .env.local
    host=localhost
    user=postgres_user
    password=postgres_password
    dbname=fufufafa_db
    port=5432
    ```
    * Ganti `localhost`, `postgres_user`, `postgres_password`, `fufufafa_db`, dan `5432` dengan detail koneksi database Anda.

3.  **Install dependencies:**
    Go akan secara otomatis mengunduh modul yang diperlukan saat Anda menjalankan atau membangun proyek. Anda bisa menjalankannya secara eksplisit jika perlu:
    ```bash
    go mod tidy
    ```

4.  **Jalankan Aplikasi:**
    ```bash
    go run main.go
    ```
    Secara default, server akan berjalan di port `5001`. Anda akan melihat output log di terminal, termasuk pesan koneksi database yang berhasil.

## Struktur Proyek

```
fufufafa-api/
├── api/
│   └── index.go        # Entry point untuk Vercel
├── config/
│   └── config.go       # Memuat konfigurasi dari .env
├── database/
│   ├── connect.go      # Logika koneksi database (GORM)
│   └── database.go     # Variabel global DB GORM
├── docs/               # File dokumentasi Swagger (JSON, YAML, Go generated)
├── handlers/
│   └── handlers.go     # Logika request handler untuk setiap endpoint
├── helpers/
│   └── helpers.go      # Fungsi bantuan (misal: mengambil data dari DB)
├── middleware/
│   └── middleware.go   # Custom middleware (misal: NotFound handler)
├── model/
│   └── model.go        # Definisi struct model data (Post)
├── public/             # Aset statis (favicon)
├── router/
│   └── router.go       # Definisi rute API
├── views/              # (Direferensikan tapi mungkin tidak digunakan aktif di Vercel build)
├── .env.local          # (File konfigurasi lokal - tidak di-commit)
├── go.mod              # Definisi modul Go
├── go.sum              # Checksum dependensi
├── main.go             # Entry point utama aplikasi (untuk running lokal)
├── swagger.json        # (Alias ke docs/swagger.json)
├── swagger.yaml        # (Alias ke docs/swagger.yaml)
└── vercel.json         # Konfigurasi deployment Vercel
```

## Endpoint API

* **`GET /`**
    * Deskripsi: Endpoint indeks, mengembalikan pesan selamat datang dan link ke dokumentasi Swagger.
    * Contoh Respon:
        ```json
        {
          "status": "Success",
          "message": "Halo User",
          "docs": "/api/docs"
        }
        ```

* **`GET /api`**
    * Deskripsi: Mengembalikan *semua* kutipan Fufufafa dari database. Endpoint ini memiliki caching.
    * Contoh Respon: `Array` dari objek `Post`.
        ```json
        [
          {
            "id": 1,
            "content": "Ini adalah kutipan pertama.",
            "datetime": "2023-10-27T10:00:00Z",
            "doksli": "Sumber Kutipan 1",
            "image_url": "http://example.com/image1.jpg"
          },
          {
            "id": 2,
            "content": "Ini adalah kutipan kedua.",
            "datetime": "2023-10-27T11:00:00Z",
            "doksli": "Sumber Kutipan 2",
            "image_url": "http://example.com/image2.jpg"
          }
          // ... kutipan lainnya
        ]
        ```

* **`GET /api/random`**
    * Deskripsi: Mengembalikan satu kutipan Fufufafa *acak* dari database.
    * Contoh Respon: Objek `Post` tunggal (acak).
        ```json
        {
          "id": 5,
          "content": "Ini adalah kutipan acak.",
          "datetime": "2023-10-27T15:00:00Z",
          "doksli": "Sumber Kutipan 5",
          "image_url": "http://example.com/image5.jpg"
        }
        ```

* **`GET /api/:quote_id`**
    * Deskripsi: Mengembalikan kutipan Fufufafa spesifik berdasarkan `id`-nya. Ganti `:quote_id` dengan ID numerik dari kutipan yang diinginkan.
    * Parameter:
        * `quote_id` (integer, path): ID unik dari kutipan.
    * Contoh Request: `GET /api/3`
    * Contoh Respon (Sukses): Objek `Post` dengan ID yang cocok.
        ```json
        {
          "id": 3,
          "content": "Ini adalah kutipan ketiga.",
          "datetime": "1403369462000", // ini adalah unix date
          "doksli": "Sumber Kutipan 3",
          "image_url": "http://example.com/image3.jpg"
        }
        ```
    * Contoh Respon (Gagal - Tidak Ditemukan):
        ```json
        {
          "error": true,
          "message": "Data tidak ditemukan"
        }
        ```
    * Contoh Respon (Gagal - Format ID Salah):
        ```json
        {
          "error": "Invalid quote ID format"
        }
        ```

## Dokumentasi API (Swagger)

Proyek ini menggunakan Swagger untuk dokumentasi API. Setelah menjalankan server secara lokal, Anda dapat mengakses UI Swagger interaktif di:

[http://localhost:5001/api/docs](http://localhost:5001/api/docs)

Di sana Anda dapat melihat detail semua endpoint, model data, dan mencoba mengirim request langsung dari browser.

> **Catatan Penting:** Dokumentasi Swagger (/api/docs) **tidak berfungsi** ketika di-deploy ke Vercel karena batasan pada serverless functions. Dokumentasi hanya tersedia saat menjalankan aplikasi secara lokal.

## Deployment

Proyek ini dikonfigurasi untuk deployment mudah di [Vercel](https://vercel.com/). File `vercel.json` mengatur rewrite rules agar semua request diarahkan ke serverless function Go (`api/index.go`). Pastikan untuk mengatur variabel lingkungan (koneksi database) di pengaturan proyek Vercel Anda.

---