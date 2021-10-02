package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type calculation struct {
	First     float64 `json:"first"`
	Second    float64 `json:"second"`
	Operation string  `json:"operation"`
}

type calcResp struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func addNumber(c echo.Context) error {
	firstNumber := c.QueryParam("first")
	secondNumber := c.QueryParam("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 64)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 64)
	summation := float64(firstNumberFloat + secondNumberFloat)
	return c.JSON(http.StatusOK, summation)
}

func deductNumber(c echo.Context) error {
	firstNumber := c.QueryParam("first")
	secondNumber := c.QueryParam("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 64)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 64)
	deduction := float64(firstNumberFloat - secondNumberFloat)
	return c.JSON(http.StatusOK, deduction)
}

func multiplyNumber(c echo.Context) error {
	firstNumber := c.QueryParam("first")
	secondNumber := c.QueryParam("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 64)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 64)
	multiplication := float64(int(firstNumberFloat * secondNumberFloat))
	return c.JSON(http.StatusOK, multiplication)
}

func divideNumber(c echo.Context) error {
	firstNumber := c.QueryParam("first")
	secondNumber := c.QueryParam("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 64)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 64)
	if secondNumberFloat == 0 {
		return c.String(http.StatusBadRequest, "Cannot divide by zero!")
	}
	division := float64(firstNumberFloat / secondNumberFloat)
	return c.JSON(http.StatusOK, division)
}

func mathFunc(c echo.Context) error {
	calc := calculation{}
	response := calcResp{}
	err := json.NewDecoder(c.Request().Body).Decode(&calc)
	firstNumber := calc.First
	secondNumber := calc.Second
	mathType := calc.Operation
	var result float64
	switch mathType {
	case "add":
		result = float64(firstNumber + secondNumber)
	case "deduct":
		result = float64(firstNumber - secondNumber)
	case "multiply":
		result = float64(firstNumber * secondNumber)
	case "divide":
		if secondNumber == 0 {
			response.Status = 1140
			response.Result = "Cannot divide by zero!"
			c.Bind(response)
			return c.JSON(http.StatusBadRequest, response)
		}
		result = float64(firstNumber / secondNumber)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	resultString := strconv.FormatFloat(result, 'f', -1, 64)
	response.Status = 200
	response.Result = resultString
	c.Bind(response)
	return c.JSON(http.StatusOK, response)
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
	e.POST("/math", mathFunc)
	e.Logger.Fatal(e.Start(":8181"))
}
