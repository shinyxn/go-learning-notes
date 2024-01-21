package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Bandung"

	// ketika kita assign ulang address2 maka address1 tidak ikut berubah, namun address2 akan membuat data baru di memory
	address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
