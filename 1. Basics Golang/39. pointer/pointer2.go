package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := &address1 // tambahkan & maka address 2 akan pointer ke address 1

	// penjabarannya sbg berikut
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // ada operator * dan &

	// data address 1 ikut berubah
	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)
}
