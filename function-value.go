package main

import "fmt"

func getGoodBye(name string) string {
	if name == "" {
		return "Hallo bro"
	} else {
		return "Hello " + name
	}

}

func main() {
	sayGoodBye := getGoodBye
	fmt.Println(sayGoodBye(""))
	fmt.Println(sayGoodBye("Ahmad Fani"))
	fmt.Println(sayGoodBye("Sopan Mukti"))
}
