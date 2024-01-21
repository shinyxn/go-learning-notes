package main

import "fmt"

func main() {

	var name1 = "Eko"
	var name2 = "Eko"

	var result bool = name1 == name2
	hasil := name1 == name2 // misal tanpa var

	fmt.Println(result)
	fmt.Println(hasil)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
