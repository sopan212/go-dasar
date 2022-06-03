package main

import "fmt"

func main() {
	//for i := 1; i <= 10; i++ {
	//	if i%2 == 1 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}

	orangs := []string{
		"Orang1",
		"orang2",
		"orang3",
	}

	for _, val := range orangs {
		if val == "orang2" {
			continue
		}
		fmt.Println(val)
	}
}
