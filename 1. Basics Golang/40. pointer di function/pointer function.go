package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address Address) {
	address.Country = "Indonesia"
}

func main() {
	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}

	ChangeAddressToIndonesia(alamat)
	fmt.Println(alamat) // alamat tidak berubah
}
