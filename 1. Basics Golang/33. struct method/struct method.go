package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// dibawah ini struct method
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuu from", a.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Indonesia"
	eko.Age = 30

	eko.sayHi("Joko")
	eko.sayHuuu()
	// fmt.Println(eko)
	// fmt.Println(eko.Name)
	// fmt.Println(eko.Address)
	// fmt.Println(eko.Age)

	// // dibawah ini menggunakan struct literals
	// joko := Customer{
	// 	Name:    "Joko",
	// 	Address: "Indonesia",
	// 	Age:     30,
	// }
	// fmt.Println(joko)

	// // dibawah harus mengikuti urutan field structnya
	// budi := Customer{"Budi", "Indonesia", 25}
	// fmt.Println(budi)
}
