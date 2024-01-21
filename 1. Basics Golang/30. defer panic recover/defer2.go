package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)
	// logging() // jika tanpa defer maka apabila error tidak akan dieksekusi
}

func main() {
	runApplication(10)
}
