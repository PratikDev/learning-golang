package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/pratikdev/learning-golang/mongoApi/router"
)

func main() {
	fmt.Println("Starting server...")

	const PORT = 3000
	portString := ":" + strconv.Itoa(PORT)
	log.Fatal(http.ListenAndServe(portString, router.Router()))

	fmt.Printf("Server running in http://localhost:%d...\n", PORT)
}
