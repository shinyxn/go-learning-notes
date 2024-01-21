package main

import "fmt"

func main() {
	var name string

	name = "Eko Kurniawan"
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	var friendName = "Budi" // tanpa deklarasi tipe data string
	fmt.Println(friendName)

	var age int8 = 30 // memaksa untuk int8
	fmt.Println(age)

	country := "Indonesia" // tidak wajib menggunakan var, namun bisa memakai := ketika deklarasi variabel
	fmt.Println(country)

	country = "America" // deklarasi selanjutnya hanya memakai =
	fmt.Println(country)
	// deklarasi multiple variable
	var (
		firstName = "Eko"
		lastName  = "Khannedy"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	var njirName string = "Brando"
	fmt.Println(njirName)

}
