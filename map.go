package main

import "fmt"

func main() {

	stockCoffe := map[string]int{
		"cofeeArabika": 10,
		"coffeBeen":    50,
		"coffeMilk":    30,
	}
	fmt.Println(stockCoffe)
	fmt.Println(stockCoffe["coffeMilk"])
	fmt.Println(stockCoffe["coffeBeen"])

	var books = make(map[string]string)
	books["One Peace"] = "Lufi"
	books["Naruto Shipuden"] = "Naruto"
	books["Boruto"] = "Boruto"

	fmt.Println(books)
	fmt.Println(len(books))
	fmt.Println(books["One Peace"])
	fmt.Println(books["Nruto Shipuden"])
	fmt.Println(books["Naruto Shipuden"])
	fmt.Println(books["Boruto"])
	delete(books, "Boruto")
	fmt.Println(books)
}
