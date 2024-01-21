package main

import "fmt"

type Man struct {
	Name string
}

// tambahkan * agar parameternya sbg pointer
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	// fmt.Println(man)
}

func main() {
	eko := Man{"Eko"}
	eko.Married()

	fmt.Println(eko) // berubah
}
