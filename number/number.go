package main

import "fmt"

func main(){
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
	//%d untuk numerik non-desimal
	//Zero value dari tipe numerik non-desimal adalah 0.

	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
	//%f untuk numerik desimal. default 6 digit
	//Zero value dari tipe numerik desimal adalah 0.0
}

// TIPE DATA NUMERIK NON DESIMAL

// Tipe Data 	| Nilai Minimum 		| Nilai Maksimum
// ----------------------------------------------------------------
// int8		| -128					| 127
// int16		| -32768				| 32767
// int32		| -2147483648			| 2147483647
// int64		| -9223372036854775808	| -9223372036854775807

// Tipe Data 	| Nilai Minimum 		| Nilai Maksimum
// ----------------------------------------------------------------
// uint8		| 0						| 255
// uint16		| 0 					| 65535
// uint32		| 0						| 4294967295
// uint64		| 0						| 18446744073709551615

// TIPE DATA NUMERIK DESIMAL

// Tipe Data 	| Nilai Minimum 		| Nilai Maksimum
// ----------------------------------------------------------------
// float32		| 1.18x10^-38			| 3.4x10^38
// float64		| 2.23x10^-308			| 1.80x10^308
// complex64	| complex numbers with float32 real & imaginary parts
// complex128	| complex numbers with float64 real & imaginary parts

// ALIAS
// Tipe Data 	| Alias untuk
// ----------------------------------------------------------------
// byte		| uint8
// rune		| int32
// int			| int32 / int64 (depend on OS)
// uint		| uint32 / uint64 (depend on OS)

// uint = unsigned integer