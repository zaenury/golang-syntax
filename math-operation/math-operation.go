package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i++ // i = i + 1
	fmt.Println(i)

	var negative = -100
	var positif = +100
	fmt.Println(negative)
	fmt.Println(positif)
}

// OPERASI MATEMATIKA

// Operator	| Keterangan
// -------------------------------------------------
// +			| Penjumlahan
// -			| Pengurangan
// *			| Perkalian
// /			| Pembagian
// %			| Sisa Pembagian

// AUGMENTED ASSIGNMENTS

// Operasi Matematika	| Augmented Assignments
// -------------------------------------------------
// a = a + 10			| a += 10
// a = a - 10			| a -= 10
// a = a * 10			| a *= 10
// a = a / 10			| a /= 10
// a = a % 10			| a %= 10

// UNARY OPERATOR

// Operator	| Keterangan
// --------------------------
// ++			| a = a + 1
// --			| a = a - 1
// -			| Negative
// +			| Positive
// !			| Boolean kebalikan
