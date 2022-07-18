package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPFarhan NoKTP = "1603070408020001"
	fmt.Println(noKTPFarhan)

	var isMarried Married = true
	fmt.Println(isMarried)
}
