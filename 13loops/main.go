package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	for _, day := range days {
		fmt.Printf("Today is %s\n", day)
	}

	value := 1

	for value < 10 {
		if value == 3 {
			goto pratik
		}

		if value == 5 {
			value++
			continue
		}

		if value == 7 {
			break
		}

		fmt.Println("value is:", value)
		value++
	}

pratik:
	fmt.Println("Name is pratik")
}
