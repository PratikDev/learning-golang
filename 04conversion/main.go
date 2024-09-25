package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What's ur age:")
	age, _ := reader.ReadString('\n')
	fmt.Println("Your age is:", age)

	ageInNum, err := strconv.ParseFloat(strings.TrimSpace(age), 32)
	if err != nil {
		panic(err)
	}

	fmt.Printf("'ageInNum' variable type is: %T\n", ageInNum)
}
