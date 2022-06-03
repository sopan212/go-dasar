package main

import "fmt"

func main() {
	var number int16 = 10000
	var conv32 int32 = int32(number)
	var con8 int8 = int8(number)

	fmt.Println(number)
	fmt.Println(conv32)
	fmt.Println(con8)

}
