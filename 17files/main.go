package main

import (
	"fmt"
	"os"
)

func createFile(filePath string, content string) (int, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}

	length, err := file.WriteString(content)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	return length, nil
}

func readFile(filePath string) (string, error) {
	byte, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(byte), nil
}

func main() {
	// Create
	length, err := createFile("./demo.txt", "something is what we wrote here")
	if err != nil {
		panic(err)
	}

	// Read
	content, err := readFile("./demo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Length for demo.txt file:", length)
	fmt.Println("Content from demo.txt file:", content)
}
