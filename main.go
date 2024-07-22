package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
}

func add[T int | float64 | string](a, b T) T {
	// a.FloatValue , aIsFloatType := a.(float64)
	return a + b
}

// 108
