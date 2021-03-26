package main

import "fmt"

func main() {
	var name = "zaenury"

	if name == "adhiim" {
		fmt.Println("Halo Adhiim")
	} else if name == `musyafa'` {
		fmt.Println(`Halo Musyafa'`)
	} else {
		fmt.Println("Halo boleh kenalan gak?")
	}

	if length := len(name); length > 5 { // short statement
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
