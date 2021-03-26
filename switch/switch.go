package main

import "fmt"

func main() {
	name := "adhiim"

	switch name {
	case "adhiim":
		fmt.Println("Halo Adhiim")
		fmt.Println("Halo Adhiim")
	case "zaenury":
		fmt.Println("Halo Zaenury")
		fmt.Println("Halo Zaenury")
	case `Musyafa'`:
		fmt.Println(`Halo Musyafa'`)
		fmt.Println(`Halo Musyafa'`)
	default:
		fmt.Println("Hai boleh kenalan gak")
		fmt.Println("Hai boleh kenalan gak")
	}

	// switch length := len(name); length > 5 { // short case
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	// versi singkatnya if pada switch

	length := len(name)
	switch { // tanpa ada variable
	case length > 10: // memanfaatkan kondisi
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
