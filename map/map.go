package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{} bisa disingkat
	person := map[string]string{
		"nama":   "zaenury",
		"alamat": "Bali",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Zaenury Adhiim"
	book["ups"] = "salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

	// FUNCTION MAP

	// Operasi						| Keterangan
	// -------------------------------------------------------
	// len(map)					| Untuk mendapat jumlah data di map
	// map[key]					| Mengambil data di map dengan key
	// map[key] = value			| Mengubah data di map dengan key
	// make(map[TypeKey]TypeValue)	| Membuat map baru
	// delete(map, key)			| Menghapus data di map dengan key
}
