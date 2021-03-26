package main

import "fmt"

func main() {
	var months = [12]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = months[4:7] // Pointernya 4, lengthnya 3, capacitynya 8
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // mendapatkan length
	fmt.Println(cap(slice1)) // mendapatkan capacity

	slice1[0] = "Mei lagi" //slice1[0] itu adalah awal array dari slice1 bukan arrayynya months
	fmt.Println(months)    //data yang ada di months kerubah

	// 3 tipe data slice = pointer, length dan capacity

	var slice2 = months[10:]
	fmt.Println(slice2)

	//KODE PROGRAM APPEND SLICE
	//adalah membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
	var slice3 = append(slice2, "Eko")
	fmt.Println(slice3)
	slice2[1] = "bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	// COPYSLICE

	copySlice := make([]string, len(slice3), cap(slice3))
	copy(copySlice, slice3)
	fmt.Println(copySlice)

	// MEMBUAT SLICE DARI ARRAY

	// Membuat Slice	| Keterangan
	// ------------------------------------------------------------------------------------------
	// array[low:high]	| Membuat slice dari array dimulai index low sampai index sebelum high
	// array[low:]		| Membuat slice dari array dimulai index low sampai index akhir array
	// array[:high]	| Membuat slice dari array dimulai dari index 0 sampai index sebelum high
	// array[:]		| Membuat slice dari array dimulai dari index 0 sampai index akhir array

	// KODE PROGRAM ARRAY VS SLICE
	// iniArray := [...]int{1. 2. 3. 4. 5}
	// iniSlice := []int{1. 2. 3. 4. 5}
}
