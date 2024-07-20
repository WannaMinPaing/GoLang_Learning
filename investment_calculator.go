package main

import "fmt"

func main() {

	for i := 0; i < 4; i++ {

		if i%2 == 0 {
			fmt.Println("Sum result", i)
		} else {
			fmt.Println("Odd result", i)
		}
	}
}
