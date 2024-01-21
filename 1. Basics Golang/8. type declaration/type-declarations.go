package main

import "fmt"

func main() {
	// type ini membuat alias tipe data baru dari tipe data yang sudah ada
	type NoKTP string // sebenarnya ini membuat alias baru dari tipe data string
	type Married bool

	var noKtpEko NoKTP = "289828420840284"
	var marriedStatus Married = true
	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)
}
