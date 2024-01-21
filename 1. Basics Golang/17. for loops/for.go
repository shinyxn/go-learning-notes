package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	// for dengan statement
	// init statement --- condition -- post statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan dengan statement ke", counter)
	}

	// mencoba mengambil data slice dengan for
	slice := []string{"Eko", "Kurniawan", "Khannedy", "Budi", "Joko"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// menggunakan for range untuk mengakses array, map, slice dsb karena lebih cepat
	for i, value := range slice {
		fmt.Println("Index ke", i, "=", value)
	}

	// menggunakan underscore _ apabila tidak ingin diwarning declared but not used
	for _, value := range slice {
		fmt.Println(value)
	}

	// apabila di map
	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
