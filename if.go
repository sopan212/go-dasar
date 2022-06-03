package main

import "fmt"

func main() {
	name := "Sopan Mukti"
	if name == "Sopan Mukti" {
		fmt.Println("Hallo " + name)
	} else {
		fmt.Println("kenalan dong..!")
	}

	//if with short statment
	if length := len(name); length > 15 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}

}
