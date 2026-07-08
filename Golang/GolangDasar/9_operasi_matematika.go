package main

import "fmt"

// Operasi Matematika
// Operator    Keterangan
// ---------------------------------------------
// +           Penjumlahan
// -           Pengurangan
// * 		   Perkalian
// /           Pembagian
// %           Sisa Pembagian

func main()  {
	var a = 10
	var b = 20
	var c = a + b
	fmt.Println(c)

	var sisa = 10 % 3
	fmt.Println(sisa)
}