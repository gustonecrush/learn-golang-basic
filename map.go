package main

import "fmt"

func main() {

	person := map[string]string{
		"name": "Farhan",
		"address": "Palembang, Sumatera Selatan",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["age"] = "19"

	book := make(map[string]string)
	book["title"] = "The end of the day"
	book["author"] = "Mark Haddan"
	book["pages"] = "89"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "pages")

	fmt.Println(book)
	fmt.Println(len(book))
}