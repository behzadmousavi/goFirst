package api

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

//Definition of body structure
type calculation struct {
	First     float64 `json:"first"`
	Second    float64 `json:"second"`
	Operation string  `json:"operation"`
	IsSave    bool    `json:"is_save"`
}

//Definition of response structure
type calcResp struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func MathFunc(c echo.Context) error {
	// DB Initializing
	db := InitDb()

	//Body structure assigning
	calc := calculation{}

	//Response structure assigning
	response := calcResp{}

	//Fetching request body
	body, _ := ioutil.ReadAll(c.Request().Body)

	//Converting body to json
	err := json.Unmarshal(body, &calc)

	//Parsing body values
	firstNumber := calc.First
	secondNumber := calc.Second
	mathType := calc.Operation
	isSave := calc.IsSave

	//Performing mathematical operations in regards to the selected operation
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

	//Filling response values
	response.Status = 200
	response.Result = result

	//Converting response to json
	json.Marshal(response)

	//Inserting the response into the DB
	if isSave == true {
		sqlStatement := "insert into math_results (first_number, second_number, result, operation) " +
			"values ($1, $2, $3, $4);"
		_, err := db.Query(sqlStatement, calc.First, calc.Second, result, calc.Operation)
		if err != nil {
			fmt.Println(err)
		}
	}

	//Returning response json
	return c.JSON(http.StatusOK, response)
}
