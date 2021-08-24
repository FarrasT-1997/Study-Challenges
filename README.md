# STUDY CHALLENGES

_API_ quiz game _untuk aplikasi Study Challenges_

# Table of Content

- [Introduction](#introduction)
- [Features](#features)
- [How to Use](#how-to-use)
    - [API Documentation](#api-documentation)

# Introduction
    Study Challenges adalah sebuah aplikasi berbasis kuis dalam bidang pendidikan, di dalam aplikasi ini seorang user dapat mengerjakan satu atau lebih set soal yang berisi 5 nomor. Satu set soal di ambil secara acak oleh sistem, namun materi pelajaran dan tingkat kesulitan dapat dipilih sesuai keinginan user.

    Untuk dapat mengakses fitur, user harus memiliki akun terlebih dahulu. Disini terdapat dua role untuk masing-masing user, antara lain : Admin dan User pengguna. Masing-masing role dapat menginisiasi soal, namun Admin memiliki wewenang khusus untuk memberikan ijin publish soal dari soal yang sudah user inisiasi sebelumnya, sehingga apabila Admin belum memberikan izin publish, maka soal dari user belum dapat di akses oleh pengguna lainnya.

### Aplikasi ini di susun dengan menggunakan bahasa/Framework :
- Golang.
- Echo Framework.
- Gorm
- Mysql


# Features

| No. | Features | Role | Keterangan |
| --- | --- | --- | --- |
| 1. | Register. | User, Admin. | |
| 2. | Login. | User, Admin. | |
| 3. | Memilih dan mengerjakan Set Soal random berdasarkan kategori dan tingkat kesulitan. | User. | Satu set berisi 5 nomor soal. |
| 4. | Melihat perolehan poin dari pengerjaan set soal. | User. | |
| 5. | Memiliki Badge yang diperoleh dari kalkulasi poin keseluruhan. | User. | Badge antara lain : Bronze, Silver, dan Gold. |
| 6. | Melihat Leaderboard. | User, Admin. | Ranking 1 s.d 10 |
| 7. | Menginisiasi soal baru, lalu dapat di akses hanya jika Admin mengizinkan _publish_. | User. | Admin dapat meng-_accept_/ me-_reject_ status soal. |
| 8. | Menginisiasi soal baru, dan langsung dapat di akses oleh setiap user. | Admin. | |
| 9. | Menerima atau menolak inisiasi soal | Admin. | |

# How to Use
## API Documentation
Berikut merupakan Endpoint yang dapat dipergunakan untuk mengakses fitur dalam aplikasi Study Challenges : [klik disini](https://app.swaggerhub.com/apis/ryanpriatama/studychallanges/1#/ "Study-challenges-endpoint")
