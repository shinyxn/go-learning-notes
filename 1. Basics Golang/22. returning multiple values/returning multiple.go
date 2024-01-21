package main

import "fmt"

// gunakan koma , untuk memisahkan antara value yg akan direturn (returning multiple value)
// jangan lupa masukkan kedalam kurung juga jika multiple, kalau 1 kan tidak
func getFullName() (string, string) {
	return "Eko", "Khannedy" // sama ini juga dipisahkan dengan koma
}

func main() {
	// bisa diassign ke multiple variable juga
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
