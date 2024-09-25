package main

import "fmt"

func main() {
	var pointer *int
	fmt.Println("Value of pointer is:", pointer)

	myName := "pratik dev"
	var myNamePointer = &myName
	fmt.Println("My name pointer is:", myNamePointer)
	fmt.Println("My name pointer value is:", *myNamePointer)
}
