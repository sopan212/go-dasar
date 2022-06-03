package main

import "fmt"

func hitungLuasSegitia(p int8, l int8) int8 {
	return p * l
}

func getHello(name string) string {
	if name == "" {
		return "hai bro!"
	} else {
		return "Hai " + name
	}
}

func main() {

	result := getHello("Sopan")
	fmt.Println(result)
	fmt.Println(getHello(""))

	hasil := hitungLuasSegitia(4, 6)
	fmt.Println(hasil)

}
