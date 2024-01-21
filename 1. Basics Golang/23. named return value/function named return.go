package main

import "fmt"

// bisa langsung dideklarasikan named valuenya
func getFullName() (firstName string, middleName string, lastName string) {
	// dapat kita assign valuenya seperti variable
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khannedy"

	return
	// langsung return semua, tanpa menyebutkan lagi satu satu macam return firstName, .., ...
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
