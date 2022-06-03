package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, val := range numbers {
		fmt.Println(val)
		total += val
	}
	return total
}

func main() {
	slices := []int{20, 40, 25, 57, 210}
	result := sumAll(10, 20, 50, 50, 100, 523)
	fmt.Println(result)
	sums := sumAll(slices...)
	fmt.Println(sums)
}
