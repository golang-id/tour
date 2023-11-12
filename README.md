# Go Tour

[![Go Reference](https://pkg.go.dev/badge/github.com/golang-id/tour.svg)](https://pkg.go.dev/github.com/golang-id/tour)

Tur Bahasa Pemrograman Go adalah sebuah situs pengenalan untuk bahasa
pemrograman Go.
Kunjungi situs https://tour.golang-id.org untuk memulai tur.

## Unduh/Pasang

Untuk memasang tur dari sumber kode, pertama
[siapkan sebuah _workspace_](https://golang.org/doc/code.html)
dan kemudian jalankan:

    $ go get github.com/golang-id/tour

Perintah tersebut akan membuat program `tour` dalam direktori `bin` di
_workspace_ anda, yang dapat dijalankan di lokal.

## Berkontribusi

Kontribusi sebaiknya mengikuti prosedur yang sama dengan proyek Go:
https://golang.org/doc/contribute.html

Untuk menjalankan server tur di lokal:
```sh
go run .
```

Peramban Anda seharusnya akan membuka halaman tur secara otomatis, jika tidak
buka langsung di [http://localhost:3999/](http://localhost:3999).

## Pelaporan Isu / Pengiriman patch

Repository ini menggunakan Github untuk perubahan kode, yang bisa diakses
lewat https://github.com/golang-id/tour/issues.

## Penerbitan

1.	Untuk menerbitkan tour.golang.org, jalankan:

	```
	gcloud --project=go-tour-id2 app deploy --no-promote app.yaml
	```
	Perintah tersebut akan membuat versi baru, yang dapat dilihat di
	[golang-org GCP project](https://console.cloud.google.com/appengine/versions?project=go-tour-id2&serviceId=default).

2.	Periksa apakah versi yang diterbitkan berjalan (klik pada tautan
	version di GCP).

3.	Jika semua berjalan dengan baik, klik "Migrate Traffic" untuk
	memindahkan 100% semua traffik ke versi yang baru.

4.	Selesai.

## Lisensi

Kecuali bila dicantumkan, berkas sumber kode go-tour didistribusikan dengan
lisensi model BSD yang bisa ditemukan pada berkas LICENSE.
