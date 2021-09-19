package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func addNumber(c echo.Context) error {
	firstNumber := c.QueryParam("first")
	secondNumber := c.QueryParam("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 32)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 32)
	var summation float32
	summation = float32(firstNumberFloat + secondNumberFloat)
	return c.JSON(http.StatusOK, summation)
}

func deductNumber(c echo.Context) error {
	firstNumber := c.QueryParam("first")
	secondNumber := c.QueryParam("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 32)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 32)
	var deduction float32
	deduction = float32(firstNumberFloat - secondNumberFloat)
	return c.JSON(http.StatusOK, deduction)
}

func multiplyNumber(c echo.Context) error {
	firstNumber := c.QueryParam("first")
	secondNumber := c.QueryParam("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 32)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 32)
	var multiplication float32
	multiplication = float32(int(firstNumberFloat * secondNumberFloat))
	return c.JSON(http.StatusOK, multiplication)
}

func divideNumber(c echo.Context) error {
	firstNumber := c.QueryParam("first")
	secondNumber := c.QueryParam("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 32)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 32)
	var division float32
	if secondNumberFloat == 0 {
		return c.String(http.StatusBadRequest, "Cannot divide by zero!")
	}
	division = float32(firstNumberFloat / secondNumberFloat)
	return c.JSON(http.StatusOK, division)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to my go playground!")
	})
	e.GET("/add", addNumber)
	e.GET("/deduct", deductNumber)
	e.GET("/multiply", multiplyNumber)
	e.GET("/divide", divideNumber)
	e.Logger.Fatal(e.Start(":8181"))
}
