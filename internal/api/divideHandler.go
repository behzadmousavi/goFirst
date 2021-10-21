package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func DivideNumber(c echo.Context) error {
	firstNumber := c.QueryParam("first")
	secondNumber := c.QueryParam("second")
	firstNumberFloat, _ := strconv.ParseFloat(firstNumber, 64)
	secondNumberFloat, _ := strconv.ParseFloat(secondNumber, 64)
	result := division(firstNumberFloat, secondNumberFloat)
	if secondNumberFloat != 0 {
		return c.JSON(http.StatusOK, result)
	} else {
		return c.JSON(http.StatusBadRequest, result)
	}
	return c.JSON(http.StatusOK, result)
}
