package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}
}

func getResponseBody(response *http.Response) string {
	bytes, err := io.ReadAll(response.Body)
	checkErrNil(err)

	responseBuilder := strings.Builder{}
	responseBuilder.Write(bytes)

	return responseBuilder.String()
}

func main() {
	// POST request (Json)
	const userJsonURL = "http://localhost:3000/user"
	performPostJsonRequest(userJsonURL)

	// POST request (Form)
	const userFormURL = "http://localhost:3000/user"
	performPostFormRequest(userFormURL)

	// GET request
	const userGetURL = "http://localhost:3000/users"
	performGetRequest(userGetURL)
}

func performGetRequest(myUrl string) {
	response, err := http.Get(myUrl)
	checkErrNil(err)
	defer response.Body.Close()

	responseBody := getResponseBody(response)

	fmt.Println(responseBody)
}

func performPostJsonRequest(myUrl string) {
	requestBody := strings.NewReader(`{"name": "pratik dev", "age":22}`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	checkErrNil(err)
	defer response.Body.Close()

	responseBody := getResponseBody(response)
	fmt.Println(responseBody)
}

func performPostFormRequest(myUrl string) {
	data := url.Values{}
	data.Add("name", "pratik dev")
	data.Add("age", "21")

	response, err := http.PostForm(myUrl, data)
	checkErrNil(err)
	defer response.Body.Close()

	responseBody := getResponseBody(response)
	fmt.Println(responseBody)
}
