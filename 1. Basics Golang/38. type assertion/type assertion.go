package main

import "fmt"

func random() interface{} {
	return "Halo"
}

func main() {
	var result interface{} = random()

	// dibawah ini contoh type assertion
	// type datanya harus sesuai dan yakin
	// jika (int) maka panic karena resultnya bukan int tapi string
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// dibawah ini contoh type assertion switch yg lebih aman
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}
