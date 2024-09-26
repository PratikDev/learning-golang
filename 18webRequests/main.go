package main

import (
	"io"
	"net/http"
	"os"
)

const URL = "https://honest-feedback.vercel.app"

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

func main() {
	response, err := http.Get(URL)
	checkErrNil(err)
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	checkErrNil(err)

	htmlContent := string(bytes)
	createFile("./index.html", htmlContent)
}

func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
