package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// dibawah ini contoh menggunakan factorial recursive (function memanggil dirinya sendiri)
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(5))
	fmt.Println(5 * 4 * 3 * 2 * 1)

	fmt.Println(factorialRecursive(5))
}
