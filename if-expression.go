package main

import "fmt"

func main() {

	name := "Aang"

	if name == "Farhan" {
		fmt.Println("Hello Farhan!")
	} else if name == "Gian" {
		fmt.Println("Hello Gian!")
	} else if name == "Aang" {
		fmt.Println("Hello Aang!")
	} else {
		fmt.Println("Hello, Boleh Kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Name is too long!")
	} else {
		fmt.Println("Name is correct!")
	}

}
