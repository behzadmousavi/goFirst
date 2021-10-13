package main

import (
	"github.com/labstack/echo/v4"
	"goFirst/app"
	"net/http"
	"strconv"
)

func divideNumber(c echo.Context) error {
	firstNumber := c.QueryParam("first")
	secondNumber := c.QueryParam("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 64)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 64)
	if secondNumberFloat == 0 {
		return c.String(http.StatusBadRequest, "Cannot divide by zero!")
	}
	division := firstNumberFloat / secondNumberFloat
	return c.JSON(http.StatusOK, division)
}
