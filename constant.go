package main

import "fmt"

func main() {
	//deklaration constan
	const name string = "Sopan Mukti"
	const hobbies = "Tracking"
	const age int8 = 28
	const married bool = true

	fmt.Println(name)
	fmt.Println(hobbies)
	fmt.Println(age)
	fmt.Println(married)

	//deklaration multiple constant

	const (
		firstName string = "sopan"
		lastName         = "Mukti"
		Age              = 28
		Hobbies          = "Tracking"
		Married   bool   = false
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(Age)
	fmt.Println(Hobbies)
	fmt.Println(Married)
}
