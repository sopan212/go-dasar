package main

import "fmt"

func main() {
	//dekalsi variable using var
	var firstName string
	firstName = "Sopan"
	fmt.Println(firstName)
	var lastName string
	lastName = "Mukti"
	fmt.Println(lastName)

	//deklarasi variabale using :=

	midleName := "Aliyansyah"
	callName := "Ganteng"
	fmt.Println(midleName)
	fmt.Println(callName)

	//multiple dekalaration variable

	var (
		friendName  = "Bagus"
		friendName2 = "Fani"
		friendName3 = "Alfian"
		friendName4 = "Hendra"
		age         = 28
		hobbies     = "Tracking"
	)
	fmt.Println(friendName)
	fmt.Println(friendName3)
	fmt.Println(friendName4)
	fmt.Println(friendName2)
	fmt.Println(age)
	fmt.Println(hobbies)

}
