package main

import "fmt"

type HashName interface {
	GetName() string
}

func sayHello(hashname HashName) {
	fmt.Println("Hello ", hashname.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
func main() {

	var sopan Person
	sopan.Name = "Sopan"

	sayHello(sopan)
	var kucing Animal
	kucing.Name = "Push"
	sayHello(kucing)
}
