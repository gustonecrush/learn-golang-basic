package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Farhan"
	names[1] = "Gian"
	names[2] = "Aang"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var ages = [3]int{
		19,
		19,
		19,
	}

	fmt.Println(ages)
	fmt.Println(len(names))
	fmt.Println(len(ages))

}