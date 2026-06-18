package main

import "fmt"

// ==========================================
// CATATAN: TYPE DECLARATIONS (Go-Lang)
// ==========================================
// * Type Declarations adalah kemampuan membuat ulang tipe data baru
//   dari tipe data yang sudah ada.
// * Type Declarations biasanya digunakan untuk membuat alias terhadap
//   tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti
// ==========================================

func main() {
	// NoKTP adalah string, NoKTP adalah nama lain dari tipe data string
	type NoKTP string

	var ktpMikhael NoKTP = "11111111"
	fmt.Println(ktpMikhael)
	fmt.Println(NoKTP("222222222"))
}
