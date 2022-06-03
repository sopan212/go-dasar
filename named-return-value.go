package main

import "fmt"

func getFullNamed2() (firstName, midleName, lastName string) {
	firstName = "Sopan"
	midleName = "Mukti"
	lastName = "Aliyansyah"

	return
}

func main() {
	a, b, c := getFullNamed2()
	fmt.Println(a, b, c)
}
