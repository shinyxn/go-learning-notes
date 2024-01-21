package main

import "fmt"

// function sebagai parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

// dibawah menggunakan type decalaration sehingga tidak repot menuliskan di parameter
// sehingga Filter adalah alias dari func(string) string
type Filter func(string) string

func sayHelloWithType(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Eko", spamFilter)
	sayHelloWithType("Anjing", spamFilter)
}
