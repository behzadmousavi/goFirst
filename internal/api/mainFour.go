package api

import "strconv"

//Defines four main arithmetic operations
func addition(a float64, b float64) string {
	result := a + b
	resultStr := strconv.FormatFloat(result, 'f', -1, 64)
	return resultStr
}

func subtraction(a float64, b float64) string {
	result := a - b
	resultStr := strconv.FormatFloat(result, 'f', -1, 64)
	return resultStr
}

func multiplication(a float64, b float64) string {
	result := a * b
	resultStr := strconv.FormatFloat(result, 'f', -1, 64)
	return resultStr
}

func division(a float64, b float64) string {
	var resultStr string
	if b != 0 {
		result := a / b
		resultStr = strconv.FormatFloat(result, 'f', -1, 64)
	} else {
		resultStr = "Cannot divide by zero!"
	}
	return resultStr
}
