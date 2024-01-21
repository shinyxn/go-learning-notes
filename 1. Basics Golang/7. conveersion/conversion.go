package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32) // konversi dari int32 ke int64
	var nilai8 int8 = int8(nilai32)    // int8 tidak akan muat untuk menghandle 100000, karena limitnya cuma 127

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Eko"
	var e = name[0] // ketika mau ngambil huruf E nya, dia tidak langsung E tapi menjadi tipe data byte/uint8
	// oleh karena itu perlu dikonversi dahulu ke string supaya bisa dapat E nya

	// konversi dari byte ke string
	var eString = string(e)
	fmt.Println(name)
	fmt.Println(e)       // byte
	fmt.Println(eString) // string
}
