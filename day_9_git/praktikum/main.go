package main

import (
	"net/http"
	"praktikum/controller"
	//"praktikum/controller"

	"github.com/labstack/echo/v4"
)

func main() {

	//initialize echo
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		result := map[string]string{
			"response_code": "200",
			"message":       "Success",
		}
		return c.JSON(http.StatusOK, result)
	})

	e.GET("/customer", controller.GetCustomer)

	e.GET("/cusomer", controller.GetCustomerId)

	e.POST("/customer", controller.CreateCustomer)

	e.PUT("/customer", controller.UpdateCustomer)

	e.DELETE("/customer", controller.DeleteCustomer)

	//define port
	e.Logger.Fatal(e.Start(":5002"))

}
