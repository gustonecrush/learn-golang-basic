package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Loop :", counter)
		counter++
	}

	fmt.Println()

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Loop :", counter)
	}

	slice := []string{"Farhan", "Gian", "Aang"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, name := range slice {
		fmt.Println(index, ".", name)
		// fmt.Println(name)
	}

	people := make(map[string]string)
	people["name"] = "Farhan"
	people["job"] = "Software Engineer"

	for key, value := range people {
		fmt.Println(key, ".", value)
	}

}
