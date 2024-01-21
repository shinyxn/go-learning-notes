package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// tambahkan * di Address agar pointer
func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}

	// lalu tambahkan & di alamat karena tipe datanya pointer
	ChangeAddressToIndonesia(&alamat)

	// atau definisikan sbg berikut
	// var alamatPointer *Address = &alamat
	// ChangeAddressToIndonesia(alamatPointer)

	fmt.Println(alamat) // alamat berubah
}
