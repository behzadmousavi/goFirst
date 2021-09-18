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
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/add", addNumber)
	myRouter.HandleFunc("/deduct", deductNumber)
	myRouter.HandleFunc("/multiply", multiplyNumber)
	myRouter.HandleFunc("/divide", divideNumber)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my Golang playground!\n")
}

func addNumber(w http.ResponseWriter, r *http.Request) {
	firstNumber := r.FormValue("first")
	secondNumber := r.FormValue("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 32)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 32)
	var summation float32
	summation = float32(firstNumberFloat + secondNumberFloat)
	fmt.Fprintln(w, summation)
}

func deductNumber(w http.ResponseWriter, r *http.Request) {
	firstNumber := r.FormValue("first")
	secondNumber := r.FormValue("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 32)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 32)
	var deduction float32
	deduction = float32(firstNumberFloat - secondNumberFloat)
	fmt.Fprintln(w, deduction)
}

func multiplyNumber(w http.ResponseWriter, r *http.Request) {
	firstNumber := r.FormValue("first")
	secondNumber := r.FormValue("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 32)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 32)
	var multiplication float32
	multiplication = float32(int(firstNumberFloat * secondNumberFloat))
	fmt.Fprintln(w, multiplication)
}

func divideNumber(w http.ResponseWriter, r *http.Request) {
	firstNumber := r.FormValue("first")
	secondNumber := r.FormValue("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 32)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 32)
	var division float32
	if secondNumberFloat == 0 {
		fmt.Fprintln(w, "Cannot divide by zero!")
	}
	division = float32(firstNumberFloat / secondNumberFloat)
	fmt.Fprintln(w, division)
}

func main() {
	fmt.Printf("Starting server at port 8080\n")
	handleRequests()
}
