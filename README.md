# Kampung GeoJSON API

Proyek ini adalah backend API yang dibangun dengan bahasa pemrograman Go untuk mengelola data GeoJSON terkait kampung atau wilayah geografis. API ini menggunakan MongoDB sebagai database untuk menyimpan dan mengambil data GeoJSON dalam format FeatureCollection.

## Fitur Utama

- **Mengelola FeatureCollection GeoJSON**: API ini memungkinkan pengguna untuk membuat, membaca, memperbarui, dan menghapus koleksi fitur GeoJSON.
- **Dukungan LineString**: Data geometri yang didukung termasuk LineString untuk merepresentasikan jalan atau batas wilayah.
- **Database MongoDB**: Data disimpan di MongoDB Atlas untuk persistensi dan skalabilitas.

## Struktur Data

Data GeoJSON disimpan dalam format FeatureCollection yang terdiri dari beberapa Feature. Setiap Feature memiliki:
- **Type**: "Feature"
- **Geometry**: Objek geometri dengan type "LineString" dan koordinat array of [longitude, latitude]
- **Properties**: Informasi tambahan seperti nama jalan, tipe jalan, dan lokasi

Contoh data GeoJSON untuk kampung:

```json
{
  "type": "FeatureCollection",
  "features": [
    {
      "type": "Feature",
      "geometry": {
        "type": "LineString",
        "coordinates": [
          [107.56278, -6.90944],
          [107.5635, -6.9102],
          [107.5642, -6.911]
        ]
      },
      "properties": {
        "name": "Jalan Kebon Kopi",
        "type": "tertiary",
        "location": "Kampung Rancabentang Kebonkopi, Kelurahan Cibeureum, Kecamatan Cimahi Selatan, Kota Cimahi"
      }
    }
  ]
}
```

## Endpoint API

- `POST /api/geo`: Membuat FeatureCollection baru
- `GET /api/geo`: Mengambil semua FeatureCollection
- `GET /api/geo/:id`: Mengambil FeatureCollection berdasarkan ID
- `PUT /api/geo/:id`: Memperbarui FeatureCollection berdasarkan ID
- `DELETE /api/geo/:id`: Menghapus FeatureCollection berdasarkan ID
- `GET /api/test-db`: Menguji koneksi database

## Teknologi yang Digunakan

- **Bahasa Pemrograman**: Go
- **Framework Web**: Fiber v2
- **Database**: MongoDB Atlas
- **Library**: MongoDB Go Driver, Godotenv untuk environment variables

## Instalasi dan Menjalankan

1. Pastikan Go sudah terinstall di sistem Anda.
2. Clone repository ini.
3. Install dependencies: `go mod tidy`
4. Buat file `.env` dengan konfigurasi database dan port:
   ```
   PORT=8080
   DB_NAME=geo
   MONGOSTRING=mongodb+srv://username:password@cluster.mongodb.net/?retryWrites=true&w=majority
   ```
5. Jalankan aplikasi: `go run main.go`

## Penggunaan

API ini dapat digunakan untuk aplikasi GIS (Geographic Information System) yang memerlukan data geografis kampung, seperti pemetaan jalan, batas wilayah, atau analisis spasial.

## Kontribusi

Untuk berkontribusi, silakan buat pull request atau issue di repository ini.