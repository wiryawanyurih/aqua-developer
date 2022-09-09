package main

import (
	"mini_project/connection"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	DB, _ := connection.Connection()

	productR := productrepository.NewProdRepository(DB)
	CategoryR := categoryrepository.NewCatRepository(DB)

	productC := productcontroller.NewProdController(productR, CategoryR)
	CategoryC := categorycontroller.NewCatController(CategoryR)

	// Product
	e.POST("/product/create", productC.CreateProduct)
	e.GET("/product/get", productC.GetProducts)
	e.GET("/product/get/:id", productC.GetProduct)
	e.PUT("/product/update/:id", productC.UpdateProduct)
	e.DELETE("/product/delete/:id", productC.DeleteProduct)

	// Category
	e.POST("/category/create", CategoryC.CreateCategory)
	e.GET("/category/get", CategoryC.GetAllCategory)
	e.GET("/category/get/:id", CategoryC.GetCategory)
	e.PUT("/category/update/:id", CategoryC.UpdateCategory)
	e.DELETE("/category/delete/:id", CategoryC.DeleteCategory)

	e.Logger.Fatal(e.Start(":5002"))
}
