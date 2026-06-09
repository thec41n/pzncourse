package main

import (
	"fmt"
)

// ==========================================
// CATATAN: KONSTANTA / CONSTANT (Go-Lang)
// ==========================================
// * Constant adalah variable yang nilainya tidak bisa diubah lagi
//   setelah pertama kali diberi nilai.
// * Cara pembuatan constant mirip dengan variable, yang membedakan
//   hanya kata kunci yang digunakan adalah 'const', bukan 'var'.
// * Saat pembuatan constant, kita wajib langsung menginisialisasikan
//   datanya.
// ==========================================
// CATATAN: DEKLARASI MULTIPLE CONSTANT (Go-Lang)
// ==========================================
// * Sama seperti variable, di Go-Lang kita juga bisa membuat constant
//   secara sekaligus banyak (multiple).
// ==========================================
func main() {
	const firstName string 	= "Devi"
	const lastName 			= "Mikhael Empi"
	const (
		middleName string 	= "Mikhael"
		age = 25
	)

	fmt.Println(firstName, middleName, lastName, age)
}
