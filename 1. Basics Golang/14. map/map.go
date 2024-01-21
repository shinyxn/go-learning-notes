package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	person["title"] = "Programmer" // misal menambahkan data lagi

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string) // membuat map tanpa assign value dahulu
	book["title"] = "Belajar Go-lang"
	book["author"] = "Eko"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book) // sesudah dihapus
	fmt.Println(len(book))

}
