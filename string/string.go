package main

import "fmt"

func main(){

	var nama string = `Zaenury Adhiim Musyafa'` 
	fmt.Printf("nama: %s \n", nama)
	//%n untuk string
	//Zero value dari string adalah "" (string kosong)

	fmt.Println(len("Zaenury"))
	fmt.Println("Zaenury Adhiim"[0])
	fmt.Println("Zaenury Adhiim"[1])

	// Function			| Keterangan
	// ---------------------------------------------------------------------
	// len("string")		| Menghitung jumlah karakter di String
	// "string"[number]	| Mengambil karakter pada posisi yang ditentukan
}