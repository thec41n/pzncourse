package main

import "fmt"

// ============================================================
// LATIHAN 2 - Kalkulator Belanja
// ============================================================
// Buat program kasir sederhana untuk menghitung total belanja.
//
// Syarat:
// 1. Buat CONSTANT untuk:
//    - Nama toko (string)
//    - Pajak dalam persen, nilainya 11 (int)
//
// 2. Buat variable untuk 3 item belanja, masing-masing punya:
//    - nama barang (string)
//    - harga (int)
//
// 3. Hitung:
//    - total harga sebelum pajak (jumlah semua harga)
//    - total pajak (total * pajak / 100)
//    - total akhir (total harga + total pajak)
//    - apakah total akhir > 200000? simpan di variable boolean
//
// 4. Tampilkan output seperti ini:
//    ===== [NAMA TOKO] =====
//    Nama toko terdiri dari [N] karakter
//    1. [nama barang 1]  : Rp [harga]
//    2. [nama barang 2]  : Rp [harga]
//    3. [nama barang 3]  : Rp [harga]
//    --------------------------
//    Subtotal            : Rp [total sebelum pajak]
//    Pajak 11%           : Rp [total pajak]
//    Total               : Rp [total akhir]
//    Lebih dari 200rb    : [true/false]
// ============================================================

func main() {
	// TODO: tulis kode kamu di sini
	const namaToko = "Toko Kelontong"
	const pajak = 11

	var (
		sabun      = "Lifebuoy"
		hargaSabun = 5000

		shampoo      = "Zinc"
		hargaShampoo = 10000

		pastaGigi      = "Pepsodent"
		hargaPastaGigi = 13000
	)

	totalHarga := (hargaPastaGigi + hargaSabun + hargaShampoo)
	totalPajak := (totalHarga * pajak / 100)
	totalAkhir := totalHarga + totalPajak
	moreThan200k := totalAkhir > 200000

	fmt.Println("=====", namaToko, "=====")
	fmt.Println("Nama Toko terdiri dari", len(namaToko), "karakter")
	fmt.Println("1. ", sabun, "  : Rp ", hargaSabun)
	fmt.Println("2. ", shampoo, ": Rp ", hargaShampoo)
	fmt.Println("3. ", pastaGigi, ": Rp ", hargaPastaGigi)
	fmt.Println("----------")
	fmt.Println("Subtotal  	 : Rp ", totalHarga)
	fmt.Println("Pajak 11% 	 : Rp ", totalPajak)
	fmt.Println("Total            : Rp ", totalAkhir)
	fmt.Println("Lebih dari 200rb : ", moreThan200k)
}
