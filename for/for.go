package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	// FOR DENGAN STATEMENT

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Zaenury", "Adhiim", `Musyafa'`}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// FOR RANGE

	for i, value := range slice {
		fmt.Println("index ke", i, "=", value)
	}
}
