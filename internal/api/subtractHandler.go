package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func SubtractNumber(c echo.Context) error {
	firstNumber := c.QueryParam("first")
	secondNumber := c.QueryParam("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 64)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 64)
	result := subtraction(firstNumberFloat, secondNumberFloat)
	return c.JSON(http.StatusOK, result)
}
