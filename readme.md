# Bioskop API 🎬

Aplikasi web sederhana untuk mengelola data bioskop menggunakan Go, Gin, dan PostgreSQL.

## Feature

- **CRUD** (**C**reate, **R**ead) data bioskop.
- **Validasi** data input untuk memastikan `Nama` dan `Lokasi` tidak kosong.
- **Basic Authentication** untuk melindungi endpoint POST.
- Koneksi ke **PostgreSQL** dengan GORM sebagai ORM.


## Setup Project

- Go 1.20+
- PostgreSQL
- Docker

### How-to

1.  **Clone repositori** atau buat folder proyek baru.

2.  **Inisialisasi Modul Go** dan instal dependensi.

    ```bash
    go mod init bioskop-api
    go get [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
    go get gorm.io/gorm
    go get gorm.io/driver/postgres
    ```

3.  **Setup Database PostgreSQL**.

    Pastikan PostgreSQL berjalan. Anda bisa menggunakan Docker:

    ```bash
    docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
    ```

    Buat database dengan nama `golang`.

4.  **Konfigurasi Koneksi Database**.

    Sesuaikan string koneksi (`dsn`) di file `database/database.go` jika diperlukan.

    ```go
    dsn := "host=localhost user=postgres password=mysecretpassword dbname=golang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
    ```

5.  **Jalankan Aplikasi**.

    ```bash
    go run main.go
    ```

    Server akan berjalan di `http://localhost:8080`.

---

## Dokumentasi API

### Basic Authentication

Untuk endpoint **POST**, Anda memerlukan kredensial Basic Auth.

- **Username**: `gaktau`
- **Password**: `bilangwow`

### Endpoint

#### 1. `POST /bioskop`

Menambahkan data bioskop baru ke database.

- **Metode**: `POST`
- **URL**: `http://localhost:8080/bioskop`
- **Authorization**: Basic Auth (lihat di atas)
- **Header**: `Content-Type: application/json`
- **Body (JSON)**:

    ```json
    {
      "nama": "XXI bla bla bla",
      "lokasi": "Somalia",
      "rating": 4.5
    }
    ```

- **Respons Sukses (201 Created)**:

    ```json
    {
        "ID": 1,
        "CreatedAt": "2023-10-26T10:00:00.000Z",
        "UpdatedAt": "2023-10-26T10:00:00.000Z",
        "nama": "XXI Gandaria City",
        "lokasi": "Jakarta",
        "rating": 4.5
    }
    ```

- **Respons Gagal (400 Bad Request)**:

    ```json
    {
        "error": "Nama dan Lokasi tidak boleh kosong"
    }
    ```

---

#### 2. `GET /bioskop`

Mengambil semua data bioskop.

- **Metode**: `GET`
- **URL**: `http://localhost:8080/bioskop`
- **Respons Sukses (200 OK)**:

    ```json
    [
      {
        "ID": 1,
        "CreatedAt": "2023-10-26T10:00:00.000Z",
        "UpdatedAt": "2023-10-26T10:00:00.000Z",
        "nama": "XXI bla bla",
        "lokasi": "Jakarta",
        "rating": 4.5
      },
      {
        "ID": 2,
        "CreatedAt": "2023-10-26T10:05:00.000Z",
        "UpdatedAt": "2023-10-26T10:05:00.000Z",
        "nama": "Cinepolis bla bla",
        "lokasi": "bogor",
        "rating": 4.8
      }
    ]
    ```

---

#### 3. `GET /bioskop/:id`

Mengambil data bioskop spesifik berdasarkan ID.

- **Metode**: `GET`
- **URL**: `http://localhost:8080/bioskop/1` (ganti `1` dengan ID bioskop yang diinginkan)
- **Respons Sukses (200 OK)**:

    ```json
    {
        "ID": 1,
        "CreatedAt": "2023-10-26T10:00:00.000Z",
        "UpdatedAt": "2023-10-26T10:00:00.000Z",
        "nama": "XXI Gandaria City",
        "lokasi": "Jakarta",
        "rating": 4.5
    }
    ```

- **Respons Gagal (404 Not Found)**:

    ```json
    {
        "error": "Bioskop tidak ditemukan"
    }
    ```