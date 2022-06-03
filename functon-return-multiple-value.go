package main

import "fmt"

func getFullName() (string, string, string) {
	return "Sopan", "Mukti", "Aliyansyah"
}

func main() {

	_, midleName, lastName := getFullName()
	fmt.Println(midleName, lastName)
}
