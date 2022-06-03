package main

import "fmt"

func sayHelloTo(firstName string, midleName string, lastName string, age int8) {
	fmt.Println("hello", firstName, midleName, lastName, age)
}

func main() {
	sayHelloTo("Sopan", "Mukti", "Aliraja", 30)
}
