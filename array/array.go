package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Zaenury"
	names[1] = "Adhiim"
	names[2] = `Musyafa'`

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//tidak bisa memanggil array yang indexnya lebih dari yang telah di deklarasikan

	var values = [3]int{
		90,
		80,
		75,
	}
	fmt.Println(values)

	fmt.Println(len(names))  //menghitung panjang array
	fmt.Println(len(values)) //menghitung panjang array
}
