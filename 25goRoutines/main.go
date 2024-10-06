package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals []string

var wg sync.WaitGroup // usally declared as pointer
var mut sync.Mutex    // usually declared as pointer

func main() {
	endPoints := []string{
		"https://google.com",
		"https://youtube.com",
		"https://github.com",
		"https://fb.com",
		"https://pratik.dev",
		"https://lco.dev",
		"https://honest-feedback.vercel.app",
	}

	for _, endPoint := range endPoints {
		wg.Add(1)
		go getStatusCode(endPoint)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endPoint string) {
	defer wg.Done()

	response, err := http.Get(endPoint)
	if err != nil {
		fmt.Println(err)
		return
	}

	mut.Lock()
	signals = append(signals, endPoint)
	mut.Unlock()

	fmt.Printf("Status code for %s is: %d\n", endPoint, response.StatusCode)
}
