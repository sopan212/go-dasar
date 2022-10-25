package main

import "fmt"

type Customer struct {
	Name, Adress string
	Age          int
	IsMaried     bool
}

//struck method
func (customer Customer) sayhei(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}
func (panggil Customer) sayLove(name string) {
	fmt.Println("Heii", name, "ilove you from ", panggil.Name)
}
func main() {

	var sopan Customer
	sopan.Name = "Sopan Mukti"
	sopan.Adress = "Tegal"
	sopan.Age = 30
	sopan.IsMaried = true
	sopan.sayLove("Vera")

	sopan.sayhei("Bagus")
	//fmt.Println(sopan.Name)
	//fmt.Println(sopan.Adress)
	//fmt.Println(sopan.Age)
	//fmt.Println(sopan.IsMaried)
	//
	//bagus := Customer{
	//	Name:     "Bagus Firmansyah",
	//	Adress:   "Bekasi",
	//	Age:      30,
	//	IsMaried: true,
	//}
	//fmt.Println(bagus.Name)
	//fmt.Println(bagus.Age)
	//fmt.Println(bagus.Adress)
	//fmt.Println(bagus.IsMaried)

}
