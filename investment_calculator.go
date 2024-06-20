package main // To know Golang where begin to start

import "fmt"

func main() {
	var number_one int = 10
	var number_two int = 10

	outputText("Hello")

	calculateValue(number_one, number_two)

	fmt.Println(returnValue(number_one, number_two))

	first_value, second_value := returnMultiValue(number_one, number_two)
	fmt.Println("The first number is ", first_value)
	fmt.Println("The second number is ", second_value)

	first_auto_value, second_auto_value := returnAutoMultiValue(number_one, number_two)
	fmt.Println("The first number is ", first_auto_value)
	fmt.Println("The second number is ", second_auto_value)

}

func outputText(text string) {
	fmt.Println(text)
}

func calculateValue(n_one int, n_two int) {
	fmt.Println(n_one + n_two)
}

func returnValue(n_one int, n_two int) int {
	return n_one + n_two
}

func returnMultiValue(n_one int, n_two int) (int, float64) {
	var number_three float64 = 10.5
	var sum_num int = n_one + n_two

	return sum_num, number_three
}

func returnAutoMultiValue(n_one int, n_two int) (sum_num int, number_three float64) {
	number_three = 10.5
	sum_num = n_one + n_two

	// return sum_num, number_three
	//same upper code
	return
}
