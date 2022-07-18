package main

import "fmt"

func main() {

	name := "Gian"

	switch name {
		case "Farhan":
			fmt.Println("Hello Farhan!")
		case "Gian":
			fmt.Println("Hello Gian!")
			fmt.Println("This is Andelle Gianzra Basae")
		default:
			fmt.Println("Hello, boleh kenalan?")
	}

	switch length := len(name); length > 5 {
		case true:
			fmt.Println("Name too length!")
		case false:
			fmt.Println("Name is correct!")
	}

	length := len(name)
	switch {
		case length > 10:
			fmt.Println("Name is too length!")
		case length < 5:
			fmt.Println("Name is not bad!")
		default:
			fmt.Println("Name has been correct!")
	}

}
