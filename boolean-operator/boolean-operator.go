package main

import "fmt"

func main(){
	var ujian = 80
	var absensi = 80

	// var lulusUjian = ujian >= 80
	// var lulusAbsensi = absensi >= 80
	// Disingkat

	var lulus = ujian >= 80 && absensi >=80 
	fmt.Println(lulus)
}

// OPERASI BOOLEAN

// Operator	| Keterangan
// ---------------------------
// &&			| Dan
// ||			| Atau
// !			| Kebalikan

// OPERASI &&

// Nilai1	| Operator	| Nilai2	| Hasil
// --------------------------------------------
// true	| &&		| true		| true
// true	| &&		| false		| false
// false	| &&		| true		| false
// false	| &&		| false		| false


// OPERASI ||

// Nilai1	| Operator	| Nilai2	| Hasil
// --------------------------------------------
// true	| ||		| true		| true
// true	| ||		| false		| true
// false	| ||		| true		| true
// false	| ||		| false		| false