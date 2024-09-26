package main

import (
	"fmt"
	"log"
)

// Function to check if a string is in a slice
func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

type ValidationResponse struct {
	FirstNumber  float32
	Operator     string
	SecondNumber float32
}

// scan and validation function
func scanAndValidation() (ValidationResponse, error) {
	OPERATORS := []string{"+", "-", "*", "/"}
	var (
		firstNumber, secondNumber float32
		operator                  string
	)

	fmt.Printf("Enter first number: ")
	// taking input for first number, also checking for errors
	if _, err := fmt.Scan(&firstNumber); err != nil {
		return ValidationResponse{}, fmt.Errorf("invalid input in first number")
	}

	fmt.Printf("Enter operator (+, -, *, /)")
	fmt.Scan(&operator)

	// checking if operator is invalid
	if !contains(OPERATORS[:], operator) {
		return ValidationResponse{}, fmt.Errorf("invalid input in operator")
	}

	fmt.Printf("Enter second number: ")
	// taking input for second number, also checking for errors
	if _, err := fmt.Scan(&secondNumber); err != nil {
		return ValidationResponse{}, fmt.Errorf("invalid input in second number")
	}

	return ValidationResponse{FirstNumber: firstNumber, Operator: operator, SecondNumber: secondNumber}, nil
}

// calculate
func calculate(response ValidationResponse) (float32, error) {
	switch response.Operator {
	case "+":
		return response.FirstNumber + response.SecondNumber, nil
	case "-":
		return response.FirstNumber - response.SecondNumber, nil
	case "*":
		return response.FirstNumber * response.SecondNumber, nil
	case "/":
		if response.SecondNumber == 0 {
			return 0, fmt.Errorf("math error: division by zero")
		}
		return response.FirstNumber / response.SecondNumber, nil
	default:
		return 0, nil
	}
}

func main() {
	response, responseErr := scanAndValidation()
	if responseErr != nil {
		log.Fatal(responseErr)
	}

	result, resultErr := calculate(response)
	if resultErr != nil {
		log.Fatal(resultErr)
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", response.FirstNumber, response.Operator, response.SecondNumber, result)
}
