package main

import "fmt"

func main() {
	const firstName string = "Eko" // constant tidak bisa diubah ubah nilainya
	const lastName = "Khannedy"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	// fmt.Println(value) // constant walau tidak dipake tidak akan error, gak seperti variable

	// Deklarasi multiple constant
	const (
		pirstName string = "Eko"
		lasName          = "Khannedy"
		palue            = 1000
	)

	fmt.Println(pirstName)
	fmt.Println(lasName)
	fmt.Println(palue)
}
