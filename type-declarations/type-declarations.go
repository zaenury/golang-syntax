package main

import "fmt"

func main(){
	type NoKTP string
	type Married bool

	var NoKtpAdhiim NoKTP = "08712340232"
	var MarriedStatus Married = false

	fmt.Println(NoKtpAdhiim)
	fmt.Println(MarriedStatus)
}