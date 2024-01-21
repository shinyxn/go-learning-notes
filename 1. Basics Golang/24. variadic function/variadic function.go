package main

import "fmt"

// ciri varargs (variadic argument) adalah dengan menambahkan ... di parameter
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	// slice sebagai parameter
	slice := []int{10, 20, 30, 40}
	total = sumAll(slice...) // tambahkan ... agar slicenya menjadi variadic
	fmt.Println(total)
}
