package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const PORT = 3000

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")

	portString := ":" + strconv.Itoa(PORT)
	log.Fatal(http.ListenAndServe(portString, router))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Pratik loves Mithila</h1>"))
}
