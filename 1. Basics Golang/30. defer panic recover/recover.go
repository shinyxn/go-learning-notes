package main

import "fmt"

func endApp() {
	// menangkap pesan error dengan recover
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("-- Aplikasinya tidak akan berhenti walaupun panic karena ada recover --")
}
