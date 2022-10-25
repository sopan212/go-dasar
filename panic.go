package main

import "fmt"

func endApp() {

	message := recover()
	if message != nil {
		fmt.Println("Error dengan message: ", message)
	}
	fmt.Print("Aplikasi Selesai ")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERORR")
	}
	fmt.Println("Aplikasi berjalan")
}
func main() {
	runApp(true)
}
