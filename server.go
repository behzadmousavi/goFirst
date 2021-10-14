package main

import (
	"github.com/behzadmousavi/goFirst/internal/api"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	// Echo Instance Initializing
	e := echo.New()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to my go playground!")
	})
	e.GET("/add", api.AddNumber)
	e.GET("/deduct", api.SubtractNumber)
	e.GET("/multiply", api.MultiplyNumber)
	e.GET("/divide", api.DivideNumber)
	e.POST("/math", api.MathFunc)
	e.Logger.Fatal(e.Start(":8181"))
}
