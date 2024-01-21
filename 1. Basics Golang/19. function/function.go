package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	sayHello()

	// misal menggunakan perulangan
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
