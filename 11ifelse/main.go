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
	fmt.Printf("Enter the login count: ")
	readedValue, _ := reader.ReadString('\n')

	loginCount, _ := strconv.ParseInt(strings.TrimSpace(readedValue), 0, 0)

	var message string

	if loginThreshold := int64(10); loginCount > loginThreshold {
		message = "You've crossed the login threshold"
	} else {
		message = "You haven't crossed the login threshold yet"
	}

	fmt.Println(message)
}
