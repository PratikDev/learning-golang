package main

import "fmt"

func adder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func main() {
	func() {
		fmt.Println("Hellow from iife")
	}()

	result := adder(2, 3, 4, 5)
	fmt.Println("Ans is:", result)
}
