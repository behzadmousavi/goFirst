package internal

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

type calculation struct {
	First     float64 `json:"first"`
	Second    float64 `json:"second"`
	Operation string  `json:"operation"`
	IsSave    bool    `json:"is_save"`
}

type calcResp struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func MathFunc(c echo.Context) error {
	db := InitDb()
	calc := calculation{}
	response := calcResp{}
	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &calc)
	firstNumber := calc.First
	secondNumber := calc.Second
	mathType := calc.Operation
	isSave := calc.IsSave
	var result string
	switch mathType {
	case "add":
		result = addition(firstNumber, secondNumber)
	case "subtract":
		result = subtraction(firstNumber, secondNumber)
	case "multiply":
		result = multiplication(firstNumber, secondNumber)
	case "divide":
		result = division(firstNumber, secondNumber)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	//resultString := strconv.FormatFloat(result, 'f', -1, 64)
	response.Status = 200
	response.Result = result
	json.Marshal(response)
	if isSave == true {
		sqlStatement := "insert into math_results (first_number, second_number, result, operation) " +
			"values ($1, $2, $3, $4);"
		_, err := db.Query(sqlStatement, calc.First, calc.Second, result, calc.Operation)
		if err != nil {
			fmt.Println(err)
		}
	}
	return c.JSON(http.StatusOK, response)
}
