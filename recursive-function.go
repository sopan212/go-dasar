package main

import "fmt"

func loopFactorial(value int) int {
	total := 1
	for i := value; i > 0; i-- {
		total *= i
	}
	return total
}

func recursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursive(value-1)
	}
}
func main() {
	result := loopFactorial(5)

	fmt.Println(result)
	//fmt.Println(5 * 4 * 3 * 2 * 1)
	hasil := recursive(5)
	fmt.Println(hasil)
}
