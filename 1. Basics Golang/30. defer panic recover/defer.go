package main

import "fmt"

func logging() {
	fmt.Println("Aplikasi selesai")
}

func runApp(isError bool) {
	// defer akan dijalankan di urutan akhir sebuah function
	// bahkan saat error sekalipun
	defer logging()
	if isError {
		panic("aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(false)
}
