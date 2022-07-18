package main

import "fmt"

func student() (string, int) {
	return "Farhan Augustiansyah", 19
}

func main() {
	name, _ := student()
	fmt.Println(name)
}
