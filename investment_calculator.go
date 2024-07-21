package main

import (
	"fmt"

	"example.com/investment-calculator/filepro"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	sum := getSum(1.1, 2.2)
	fmt.Println(sum)
	tellSomething()

	message := filepro.ShowMessage("I am custom package")

	fmt.Println(message)
	fmt.Println("Main File")
	fmt.Println(randomdata.Email())
}
