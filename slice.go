package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Julli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	fmt.Println(months)
	var slice = months[6:9]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	newSlice := make([]string, 6, 9)

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	newSlice[0] = "apel"
	newSlice[1] = "jeruk"
	newSlice[2] = "mangga"
	newSlice[3] = "alpuket"
	newSlice[4] = "nanas"
	newSlice[5] = "jambu"

	fmt.Println(newSlice)
}
