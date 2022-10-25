package main

import "fmt"

func logging() {
	fmt.Print("function selesai di eksekusi")
}

func runAplication(value int) {
	defer logging()
	result := 10 / value
	fmt.Println(result)
	fmt.Println("Run aplication")
}
func main() {
	runAplication(2)

}
