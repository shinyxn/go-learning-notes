package main

import "fmt"

// masukkan parameter dan tipe datanya
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Eko", "Khannedy")
}
