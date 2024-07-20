package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getNumberFromFile() (float64, error) {
	data, err := os.ReadFile("wrong.txt")

	if err != nil {
		return 0, errors.New("Fail to read file.")
	}

	numberText := string(data)
	number, err := strconv.ParseFloat(numberText, 64)

	if err != nil {
		return 0, errors.New("Failed tp parse stored balance value.")
	}

	//strconv => string to conversition
	// float 64
	return number, nil
}

func writeNumberToFile(number int) {
	numberText := fmt.Sprint(number)
	os.WriteFile("number.txt", []byte(numberText), 0644)
}

func main() {
	writeNumberToFile(2)

	var getNumber, err = getNumberFromFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}

	fmt.Println(getNumber)
}
