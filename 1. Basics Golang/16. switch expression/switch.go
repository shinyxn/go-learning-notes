package main

import "fmt"

func main() {
	name := "Joko"

	switch name {
	case "Eko":
		fmt.Println("Halo Eko")
	case "Joko":
		fmt.Println("Halo Joko")
	default:
		fmt.Println("Hai, kenalan dong")
	}

	// switch dengan short statement

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
