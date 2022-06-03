package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFilter := filter(name)
	fmt.Println("Hello " + nameFilter)
}

func spamFiltered(name string) string {
	switch name {
	case "Anjing":
		{
			return "..."
		}
	case "Babi":
		{
			return "..."
		}
	case "Kunyuk":
		{
			return "..."
		}
	case "Ngentot":
		{
			return "..."
		}
	default:
		return name
	}
}

func main() {
	sayHelloWithFilter("Span Mukti", spamFiltered)
	sayHelloWithFilter("Anjing", spamFiltered)
	sayHelloWithFilter("Ngentot", spamFiltered)
}
