package main

import "fmt"

func main() {
	type DataKtp string
	type Married bool

	var noKtpSopan DataKtp = "31730302132131"
	var nama DataKtp = "Sopan Mukti"
	var marriedStatus Married = true

	fmt.Println(noKtpSopan)
	fmt.Println(nama)
	fmt.Println(marriedStatus)
}
