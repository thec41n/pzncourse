package main

import "fmt"

// ============================================================
// LATIHAN 1 - Biodata Digital
// ============================================================
// Buat program yang menyimpan dan menampilkan biodata seseorang.
//
// Syarat:
// 1. Buat variable untuk menyimpan data berikut:
//    - nama       : nama lengkap kamu (string)
//    - umur       : umur kamu sekarang (int)
//    - tinggi     : tinggi badan dalam cm (float64)
//    - kota       : kota asal kamu (string)
//    - isPelajar  : apakah kamu masih pelajar/mahasiswa? (bool)
//
// 2. Tampilkan semua data dengan format seperti ini:
//    ===== BIODATA SAYA =====
//    Nama     : [nama]
//    Umur     : [umur] tahun
//    Tinggi   : [tinggi] cm
//    Kota     : [kota]
//    Pelajar  : [true/false]
//    Panjang nama: [jumlah huruf] karakter
//
// 3. Gunakan minimal SATU var block (var(...)) untuk deklarasi variable
// 4. Gunakan minimal SATU shorthand := untuk salah satu variable
// ============================================================

func main() {
	// TODO: tulis kode kamu di sini
	nama := "Joko Noto Boto"
	var (
		umur              = 24
		tinggi    float32 = 172.5
		kota      string  = "Jakarta"
		isPelajar bool    = true
	)

	fmt.Println("======== BIODATA SAYA ========")
	fmt.Println("Nama               : ", nama)
	fmt.Println("Umur               : ", umur, "tahun")
	fmt.Println("Tinggi             : ", tinggi, "cm")
	fmt.Println("Kota               : ", kota)
	fmt.Println("Pelajar            : ", isPelajar)
	fmt.Println("Panjang Nama       : ", len(nama))
}
