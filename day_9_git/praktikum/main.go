package main

import (
	"day_7_git/controller"
	"net/http"

	"github.com/labstack/echo/v4"
	//"gorm.io/driver/postgres"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

func main() {
	// dsn := "host=localhost user=postgres password=AquaYurih123 dbname=efishery_academy port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
