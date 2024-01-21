package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1
	// bahkan address3 juga ikut berubah karena mengacu pada memori yg sama
	// adanya operator * pada address 2 di line 18 membuat semua akan refer ke memori yang sama

	address2.City = "Bandung"

	// dengan menambahkan * maka address1 akan ikut refer ke memori yang sama pada address2 yg baru diassign.
	// bintang dibawah nih, & nya juga gak ada, bukan spt file sebelumnya
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// membuat pointer baru dengan function new
	var address4 *Address = new(Address)
	address4.City = "Bogor"

	// dibawah ini adalah penjabarannya apabila tanpa new
	// var address4 *Address = &Address{"...", "...", "..."}

	fmt.Println(address4)
}
