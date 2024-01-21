package main

import "fmt"

type Blacklist func(name string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome...", name)
	}
}

// dibawah ini kalau make function as parameter biasanya
func isBlack(name string) bool {
	if name == "admin" {
		return true
	} else {
		return false
	}
}

// return name == kan true
func isRoot(name string) bool {
	return name == "root"
}

func main() {
	registerUser("admin", isBlack)
	registerUser("root", isRoot)

	// dibawah ini contoh menggunakan anonymous function (yg disimpan dalam variable)
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)

	// dibawah ini anonymous function tapi langsung jadi parameter
	registerUser("eko", func(name string) bool {
		return name == "user"
	})

}
