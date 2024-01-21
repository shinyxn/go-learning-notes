package main

import "fmt"

// function sebagai value
func getGoodBye(name string) string {
	return "Goodbye " + name
}

func main() {
	// variable ini memanggil function yg sebagai value, tanpa tanda kurung ()
	goodbye := getGoodBye
	fmt.Println(goodbye("Eko")) // lalu variable nya bisa digunakan layaknya function

}
