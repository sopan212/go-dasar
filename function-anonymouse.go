package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You Are Blacklist", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blaclistAdmin := func(name string) bool {
		return name == "Admin"
	}
	registerUser("Admin", blaclistAdmin)
	registerUser("Sopan", blaclistAdmin)

}
