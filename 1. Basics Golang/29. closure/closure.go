package main

import "fmt"

func main() {
	name := "Eko"
	counter := 0

	increment := func() {
		name = "Budi" // scope name akan dirubah jadi budi
		//name := "Yanto" // membuat variable name sendiri yg terpisah dari scope diatasnya
		fmt.Println("Increment")
		fmt.Println(name)
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
