package main

import "fmt"

func sayGreeter(name string, age int8, greeter func(string, int8) string) {
	greetResult := greeter(name, age)
	fmt.Println("Halo", greetResult)
}

func greetFilter(name string, age int8) string {
	if name == "Chitanda" && age == 15 {
		return "Yang mulia"
	} else if age > 15 {
		return "Rakyat tua jelata"
	} else {
		return "Rakyat kecil jelata"
	}
}

// .................
// misalkan guna type
type Greeter func(string, int8) string

func sayGreetBro(name string, age int8, greeter Greeter) {
	greetOne, greetTwo := greetFilterFinal(name, age)
	fmt.Println("Halo", greetOne, greetTwo)
}

func greetFilterFinal(name string, age int8) (string, string) {
	if name == "Chitanda" {
		switch {
		case age > 20:
			return "Yang mulia", "anda sudah tua ya"
		case age > 15:
			return "Yang mulia", "semangat menjalani hari"
		default:
			return "Adinda", "mau makan apa hari ini?"
		}
	} else {
		switch {
		case age > 20:
			return name, "anda sudah tua ya"
		case age > 15:
			return name, "semangat menjalani hari"
		default:
			return name, "mau makan apa hari ini?"
		}
	}
}

func main() {
	// sayGreeter("Bukan chitanda", 16, greetFilter)
	sayGreetBro("Chitanda", 21, greetFilter)
	sayGreetBro("Arisu", 10, greetFilter)
}
