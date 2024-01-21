package main

import "fmt"

// didalam kontrak HasName mempunyai method yaitu getName()
type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var eko Person
	eko.Name = "Eko"

	cat := Animal{
		Name: "Push",
	}

	sayHello(eko)
	sayHello(cat)

}
