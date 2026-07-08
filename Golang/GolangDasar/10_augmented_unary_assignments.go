package main

import "fmt"

// Augmented Assignments
// 
// Augmented Assignments (atau compound assignment operators) adalah cara singkat 
// untuk melakukan operasi matematika pada sebuah variabel dan langsung memasukkan 
// hasilnya kembali ke variabel itu sendiri (singkatan dari operasi biasa).
//
// Operasi Matematika    Augmented Assignments
// ---------------------------------------------
// a = a + 10            a += 10
// a = a - 10            a -= 10
// a = a * 10            a *= 10
// a = a / 10            a /= 10
// a = a % 10            a %= 10

// Unary Operator
// 
// Unary Operator adalah operator yang hanya membutuhkan SATU nilai atau 
// SATU variabel saja untuk melakukan suatu operasi (berbeda dengan operasi 
// biasa seperti 'a + b' yang membutuhkan dua variabel).
//
// Operator    Keterangan & Contoh Mudah
// -------------------------------------------------------------------------
// ++          Menambah nilai variabel sebanyak 1 (Contoh: jika a=5, jadi 6)
// --          Mengurangi nilai variabel sebanyak 1 (Contoh: jika a=5, jadi 4)
// -           Mengubah angka menjadi negatif (Contoh: 5 menjadi -5)
// +           Menandakan angka positif (Contoh: +5, biasanya tidak wajib ditulis)
// !           Membalikkan nilai boolean/logika (Contoh: benar menjadi salah)

func main()  {

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	// Unary Operator
	var j = 1
	var k = 5
	j++ // j = j+1
	k-- // k = k-1
	fmt.Println(j)
	fmt.Println(k)
}