package main

import "fmt"

// Operasi Perbandingan
// 
// Operasi Perbandingan digunakan untuk membandingkan dua buah data.
// Hasil dari operasi ini selalu berupa nilai boolean, yaitu menentukan 
// apakah perbandingan tersebut benar atau salah.
//
// Kondisi        Hasil Nilai
// --------------------------
// Benar          true
// Salah          false
// ---
// Operator Perbandingan
// 
// Operator perbandingan adalah simbol-simbol yang digunakan untuk 
// membandingkan dua buah data atau variabel.
//
// Operator    Keterangan
// ---------------------------------------------
// >           Lebih Dari
// <           Kurang Dari
// >=          Lebih Dari Sama Dengan
// <=          Kurang Dari Sama Dengan
// ==          Sama Dengan
// !=          Tidak Sama Dengan

func main()  {
	var name1 = "Joko"
	var name2 = "Joko"

	var result bool = name1 != name2
	fmt.Println(result)
}