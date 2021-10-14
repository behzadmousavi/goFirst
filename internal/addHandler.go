package internal

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func AddNumber(c echo.Context) error {
	firstNumber := c.QueryParam("first")
	secondNumber := c.QueryParam("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 64)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 64)
	result := addition(firstNumberFloat, secondNumberFloat)
	return c.JSON(http.StatusOK, result)
}
