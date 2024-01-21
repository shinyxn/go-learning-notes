package main

import "fmt"

func main() {

	// contoh kode break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke", i)
	}

	// contoh continue
	for i := 0; i < 10; i++ {

		//dibawah ini artinya jika habis dibagi 2 (bilangan genap) maka akan dicontinue/diskip gak diprint
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke", i)
	}
}
