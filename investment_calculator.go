package main

import "fmt"

func main() {
	age := 32
	var agePointer *int
	agePointer = &age

	fmt.Println(age)
	getAdultYear(agePointer)
	fmt.Println(age)
}

func getAdultYear(agePointerRef *int) {
	*agePointerRef = *agePointerRef - 18
}
