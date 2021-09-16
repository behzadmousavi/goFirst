package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", nil)
	myRouter.HandleFunc("/add", addNumber)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func addNumber(w http.ResponseWriter, r *http.Request) {
	first_number := r.FormValue("first")
	second_number := r.FormValue("second")
	first_number_int,_ := strconv.Atoi(first_number)
	second_number_int,_ := strconv.Atoi(second_number)
	var summation int
	summation = first_number_int + second_number_int
	fmt.Fprintln(w, summation)
}

func main() {
	fmt.Printf("Starting server at port 8080\n")
	handleRequests()
}