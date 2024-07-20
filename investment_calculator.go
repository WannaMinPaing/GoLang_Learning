package main

import (
	"fmt"
	"os"
	"strconv"
)

func getNumberFromFile() float64 {
	data, _ := os.ReadFile("number.txt")
	numberText := string(data)
	number, _ := strconv.ParseFloat(numberText, 64)
	//strconv => string to conversition
	// float 64
	return number
}

func writeNumberToFile(number int) {
	numberText := fmt.Sprint(number)
	os.WriteFile("number.txt", []byte(numberText), 0644)
}

func main() {
	writeNumberToFile(2)
	var getNumber = getNumberFromFile()
	fmt.Println(getNumber)
}
