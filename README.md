# Tur bahasa pemrograman Go

Tur Bahasa Pemrograman Go adalah sebuah situs pengenalan untuk bahasa
pemrograman Go.
Kunjungi situs
<https://tour.golang-id.org>
untuk memulai tur.


## Pasang

Untuk memasang dan menjalankan tur dari sumber kode, memerlukan
perkakas
[Git SCM](https://git-scm.com/)
dan
[Go](https://golang-id.org/doc/install/).

Pertama, salin repositori ini ke lokal,

```
$ git clone https://github.com/golang-id/tour.git
```

Kemudian, jalankan server Tour di lokal:

```
$ go run .
```

Sebuah peramban akan membuka halaman tur secara otomatis, jika
tidak, buka langsung di
<http://localhost:3999/>


## Pelaporan isu dan perbaikan

Repository ini menggunakan GitHub untuk perubahan kode.

Isu dapat diakses dan dilaporkan lewat
<https://github.com/golang-id/tour/issues>.

Perubahan dapat dikirim lewat _pull request_ ke
<https://github.com/golang-id/tour/pulls>.


## Penerbitan

(Bagian ini hanya khusus untuk pengelola situs).

Untuk menerbitkan versi Tour yang baru ke tour.golang-id.org, jalankan:

```
$ gcloud --project=go-tour-id2 app deploy --no-promote app.yaml
```

Perintah tersebut akan membuat versi baru, yang dapat dilihat di
[golang-org GCP project](https://console.cloud.google.com/appengine/versions?project=go-tour-id2&serviceId=default).
Periksa apakah versi yang baru berjalan dengan meng-klik pada tautan
"version" di GCP.
Jika semua berjalan dengan baik, klik "Migrate Traffic" untuk
memindahkan semua koneksi ke versi yang baru.


## Lisensi

Berkas sumber kode go-tour di-distribusi dengan lisensi model BSD yang bisa
ditemukan pada berkas LICENSE, kecuali bila dicantumkan berbeda.
