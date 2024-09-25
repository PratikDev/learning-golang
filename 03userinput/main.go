package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What's ur age: ")

	input, _ := reader.ReadString('\n')

	fmt.Println("Your age is:", input)
}
