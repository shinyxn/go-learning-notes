package main

import "fmt"

// masukkan tipe data pengembaliannya supaya dapat return value. func(...) tipedatanya {}
func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func main() {
	// mengambil data kembaliannya
	result := getHello("Eko")
	fmt.Println(result)

	// atau langsung tanpa buat variabel
	fmt.Println(getHello(""))
}
