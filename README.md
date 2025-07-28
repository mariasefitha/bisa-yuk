# 🫶 Donasi Yuk! — Final Project Golang

"Donasi Yuk!" adalah aplikasi backend REST API yang dibuat dengan bahasa Go (Golang), dirancang untuk memfasilitasi kampanye donasi secara online. Aplikasi ini mendukung proses registrasi pengguna, pembuatan kampanye donasi, serta fitur donasi dari pengguna lain.

---

## 🚀 Fitur-Fitur Utama

### 🔐 Autentikasi
- **Register** (`POST /api/users/register`)
- **Login** (`POST /api/users/login`)
- Menghasilkan **JWT Token** yang digunakan untuk mengakses endpoint yang dilindungi.

### 📣 Manajemen Kampanye Donasi
- **Create Campaign** (`POST /api/campaigns/`)
- **Get All Campaigns** (`GET /api/campaigns/`)
- **Get Campaign Detail** (`GET /api/campaigns/:id`)
- **Update Campaign** (`PUT /api/campaigns/:id`)
- **Delete Campaign** (`DELETE /api/campaigns/:id`)

> Semua operasi kampanye membutuhkan autentikasi (token JWT).

### 💰 Manajemen Donasi
- **Create Donasi** (`POST /api/donasi/`)
- **Get Donasi by User** (`GET /api/donasi/user`)
- **Get Donasi by Campaign** (`GET /api/donasi/campaign/:id`)
- **Get All Donasi** (`GET /api/donasi/`)

> Semua endpoint donasi juga memerlukan autentikasi JWT.

---

## 🛠️ Teknologi yang Digunakan

- **Go (Golang)** dengan [Gin](https://github.com/gin-gonic/gin) framework
- **PostgreSQL** via [GORM](https://gorm.io/)
- **JWT Authentication** untuk proteksi endpoint
- **Railway** untuk deployment
- **Postman** untuk testing endpoint

---

## ⚙️ Instalasi Lokal

### 1. Clone Repository
bash
git clone https://github.com/username/donasi-yuk.git
cd donasi-yuk


### 2. Konfigurasi .env
Buat file .env dan tambahkan:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=donasi_yuk_db
JWT_SECRET=secret123


### 3. Jalankan Aplikasi
go mod tidy
go run main.go

---

☁️ Deploy ke Railway
Buat project baru di Railway
Tambahkan plugin PostgreSQL
Isi Environment Variable seperti:
DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME
JWT_SECRET
Hubungkan dengan GitHub dan deploy otomatis dari branch master

---

# POSTMAN:
#####Ini contoh#####

1. POST https://bisa-yuk-production.up.railway.app/api/users/register
-tambahkan di Body:
{
  "name": "Lilis",
  "email": "lilis@mail.comm",
  "password": "lilis123"
}
-tambahkan di Scripts (Post-response)
const res = pm.response.json();
if (res.token) {
    pm.environment.set("token", res.token);
    console.log("✅ Token berhasil disimpan ke environment!");
} else {
    console.warn("❌ Tidak ada token ditemukan di respons!");
}

2. POST https://bisa-yuk-production.up.railway.app/api/users/login
-tambahkan di Body:
{
  "email": "lilis@mail.comm",
  "password": "lilis123"
}
-tambahkan di Scripts (Post-response)
const res = pm.response.json();
if (res.token) {
    pm.collectionVariables.set("token-login", res.token);
    console.log("✅ Token berhasil disimpan ke environment!");
} else {
    console.warn("❌ Tidak ada token ditemukan di respons!");
}

3. POST https://bisa-yuk-production.up.railway.app/api/campaigns/
-tambahkan Authorization
Type: Bearer Token
Token: {{token-login}}
-tambahkan Headers
key: X-LOGIN
value: {{token-login}}
-tambahkan Body
{
  "title": "Ayo Bantu Berobat Anak Stunting",
  "description": "Anak Stunting harus kita obati dan kita jaga, mari dukung untuk biaya perobatan anak-anak stunting dan cegah anak agar tidak stunting!",
  "target_amount": 2000000000
}

---

📌 Catatan
Pastikan semua endpoint yang dilindungi diberikan header:
Authorization: Bearer <your-token>
