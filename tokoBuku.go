package main

import "fmt"

func main() {
	var jumlah_buku int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&jumlah_buku)

	// Hitung total eksemplar
	jumlah_eksemplar := jumlah_buku * 10
	harga_per_eksemplar := 5000

	// Fungsi untuk menghitung total harga dengan diskon
	hitungHarga := func(jumlah int, diskon int) int {
		return jumlah * harga_per_eksemplar * (100 - diskon) / 100
	}

	var total_harga int

	if jumlah_eksemplar <= 100 {
		total_harga = hitungHarga(jumlah_eksemplar, 0)
	} else if jumlah_eksemplar <= 200 {
		harga_100 := hitungHarga(100, 5)
		harga_sisa := hitungHarga(jumlah_eksemplar-100, 15)
		total_harga = harga_100 + harga_sisa
	} else {
		harga_100 := hitungHarga(100, 7)
		harga_100_kedua := hitungHarga(100, 17)
		harga_sisa := hitungHarga(jumlah_eksemplar-200, 27)
		total_harga = harga_100 + harga_100_kedua + harga_sisa
	}

	fmt.Printf("Total harga yang harus dibayar: Rp.%d", total_harga)
}
