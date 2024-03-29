package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	// var contohError error = errors.New("Ups error")
	// fmt.Println(contohError.Error())

	// contoh implementasi pengecekan error
	hasil, err := Pembagian(10, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
