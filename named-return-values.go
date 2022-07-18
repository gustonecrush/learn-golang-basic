package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Farhan"
	middleName = "Augustiansyah"
	lastName = "Tobing"

	return
}

func main() {
	firstName, middleName, _ := getCompleteName() // penamaan variable sebenarnya bebas gak harus sama dengan named returnnya
	fmt.Println(firstName, middleName)
}