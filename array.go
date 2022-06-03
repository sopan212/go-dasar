package main

import "fmt"

func main() {
	var fruits [10]string

	fruits[0] = "apel"
	fruits[1] = "nanas"
	fruits[2] = "semangka"
	fruits[3] = "jeruk"
	fruits[4] = "salak"
	fruits[5] = "timun"
	fruits[6] = "nangka"
	fruits[7] = "jambu"
	fruits[8] = "cery"
	fruits[9] = "pir"

	fmt.Println(fruits)
	fmt.Println(fruits[0])
	fmt.Println(fruits[1])
	fmt.Println(fruits[2])
	fmt.Println(fruits[3])
	fmt.Println(fruits[4])
	fmt.Println(fruits[5])
	fmt.Println(fruits[6])
	fmt.Println(fruits[7])
	fmt.Println(fruits[8])
	fmt.Println(fruits[9])

	var values = [4]int{
		3,
		4,
		5,
	}

	fmt.Println(values)

	var angkaLagi = [5]int{
		30, 40, 100, 400,
	}
	fmt.Println(angkaLagi)
}
