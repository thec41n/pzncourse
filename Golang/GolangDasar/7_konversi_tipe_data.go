package main

import "fmt"

// ==========================================
// CATATAN: KONVERSI TIPE DATA (Go-Lang)
// ==========================================
// * Di Go-Lang kadang kita butuh melakukan konversi tipe data dari
//   satu tipe ke tipe lain.
// * Misal kita ingin mengkonversi tipe data int32 ke int64, dan lain-lain.
// ==========================================

func main()  {
	var nilai32 int32 = 40000
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Joko Noto Boto"
	// dari uint8
	var e = name[1]
	// jadi string lagi
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(eString)
	fmt.Println(e)
}