package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khannedy"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		95,
		90,
		80,
	}

	fmt.Println(values)
	fmt.Println(len(names)) // untuk melihat jumlah panjang array
	fmt.Println(len(values))

}
