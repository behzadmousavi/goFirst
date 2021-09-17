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
	myRouter.HandleFunc("/deduct", deductNumber)
	myRouter.HandleFunc("/multiply", multiplyNumber)
	myRouter.HandleFunc("/divide", divideNumber)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func addNumber(w http.ResponseWriter, r *http.Request) {
	firstNumber := r.FormValue("first")
	secondNumber := r.FormValue("second")
	firstNumberInt, _ := strconv.Atoi(firstNumber)
	secondNumberInt, _ := strconv.Atoi(secondNumber)
	var summation int
	summation = firstNumberInt + secondNumberInt
	fmt.Fprintln(w, summation)
}

func deductNumber(w http.ResponseWriter, r *http.Request) {
	firstNumber := r.FormValue("first")
	secondNumber := r.FormValue("second")
	firstNumberInt, _ := strconv.Atoi(firstNumber)
	secondNumberInt, _ := strconv.Atoi(secondNumber)
	var deduction int
	deduction = firstNumberInt - secondNumberInt
	fmt.Fprintln(w, deduction)
}

func multiplyNumber(w http.ResponseWriter, r *http.Request) {
	firstNumber := r.FormValue("first")
	secondNumber := r.FormValue("second")
	firstNumberInt, _ := strconv.Atoi(firstNumber)
	secondNumberInt, _ := strconv.Atoi(secondNumber)
	var multiplication int
	multiplication = firstNumberInt * secondNumberInt
	fmt.Fprintln(w, multiplication)
}

func divideNumber(w http.ResponseWriter, r *http.Request) {
	firstNumber := r.FormValue("first")
	secondNumber := r.FormValue("second")
	firstNumberInt, _ := strconv.Atoi(firstNumber)
	secondNumberInt, _ := strconv.Atoi(secondNumber)
	var division int
	division = firstNumberInt / secondNumberInt
	fmt.Fprintln(w, division)
}

func main() {
	fmt.Printf("Starting server at port 8080\n")
	handleRequests()
}
