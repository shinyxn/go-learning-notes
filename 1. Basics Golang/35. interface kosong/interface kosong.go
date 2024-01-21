package main

import "fmt"

type Apapun interface{}

// btw di golang terbaru ada any
// interface kosong = kontrak tipe data apapun
// interface kosong juga bisa ditaruh sebagai parameter di function cth func Ups(i int, apapun interface{}) ...
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func Njir(i int) Apapun {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Njir"
	}
}

func main() {
	// dibawah ini tidak bisa, karena return datanya bukan int tapi interface kosong
	// var data int = Ups(1)

	var data interface{} = Ups(1)
	fmt.Println(data)

	result := Ups(2)
	fmt.Println(result)

	hasil := Njir(3)
	fmt.Println(hasil)
}
