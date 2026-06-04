package main

import "fmt"

// ==========================================
// CATATAN: VARIABLE (Go-Lang)
// ==========================================
// * Variable adalah tempat untuk menyimpan data.
// * Variable digunakan agar kita bisa mengakses data yang sama
//   dimanapun kita mau.
// * Di Go-Lang, Variable hanya bisa menyimpan tipe data yang sama.
//   Jika kita ingin menyimpan data yang berbeda-beda jenis,
//   kita harus membuat beberapa variable.
// * Untuk membuat variable, kita bisa menggunakan kata kunci 'var',
//   lalu diikuti dengan nama variable dan tipe datanya.
// ==========================================
// CATATAN: TIPE DATA VARIABLE (Go-Lang)
// ==========================================
// * Saat kita membuat variable, maka kita wajib menyebutkan
//   tipe data variable tersebut.
// * Namun jika kita langsung menginisialisasikan data pada
//   variable nya, maka kita tidak wajib menyebutkan tipe
//   data variable nya.
// ==========================================
// CATATAN: KATA KUNCI VAR (Go-Lang)
// ==========================================
// * Di Go-Lang, kata kunci var saat membuat variable tidak lah wajib.
// * Asalkan saat membuat variable kita langsung menginisialisasi datanya.
// * Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan
//   kata kunci := saat menginisialisasikan data pada variable tersebut.
// ==========================================
func main() {
	var name string

	name = "Joko Noto Boto"
	fmt.Println(name)

	name = "Budi Joko"
	fmt.Println(name)

	var angka = 1
	fmt.Println(angka)

	nama := "Jokoo"
	fmt.Println(nama)
}
