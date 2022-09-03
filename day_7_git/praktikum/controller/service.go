package controller

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

var CustomerList = make(map[int]Customer)

var LastID int = 1

func GetCustomer(c echo.Context) (err error) {
	var result []Customer

	// dataDummyCustomer := Customer{
	// 	ID:    1,
	// 	name:  "Yurih",
	// 	Email: "yurih@mail.com",
	// }

	result = append(result, dataDummyCustomer)

	for key, _ := range CustomerList {
		tempCustomer := CustomerList[key]

		result = append(result, tempCustomer)
	}

	return c.JSON(http.StatusOK, result)
}

func CreateCustomer(c echo.Context) error {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	log.Println("Masuk ke create user")
	req := new(Customer)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	log.Println("Req data masuk", req)
	user := Customer{
		ID:    LastID,
		name:  req.name,
		Email: req.Email,
		//DescriptionUser : req.DescriptionUser
	}

	CustomerList[LastID] = user
	LastID++

	return c.JSON(http.StatusCreated, user)
}

func GetCustomerId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Println("ID Customer: ", id)
	return c.JSON(http.StatusOK, CustomerList[id])
}

func UpdateCustomer(c echo.Context) error {
	cust := new(Customer)
	var AllCust = []Customer{}
	if err := c.Bind(cust); err != nil {
		return err
	}

	Id, _ := strconv.Atoi(c.Param("id"))
	UpCust := Customer{
		ID:    AllCust[Id].ID,
		name:  cust.name,
		Email: cust.Email,
	}

	AllCust[Id] = UpCust

	return c.JSON(http.StatusOK, AllCust[Id])
}

func DeleteCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(CustomerList, id)
	log.Println("Customer terhapus")
	return c.NoContent(http.StatusNoContent)
}
