package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"systementor.se/goapi0124/data"
)

var employees []data.Employee

func apiEmployee(c *gin.Context) {
	//var employees []data.Employee
	//data.DB.Find(&employees)

	c.IndentedJSON(http.StatusOK, employees)
}

func apiEmployeeById(c *gin.Context) {
	id := c.Param("id")

	for _, emp := range employees {
		i, _ := strconv.Atoi(id)
		if emp.Id == i {
			c.JSON(http.StatusOK, emp)
		}
	}

	c.String(http.StatusNotFound, "")
}

func main() {
	employees = append(employees, data.Employee{
		Id:   1,
		Age:  15,
		City: "Stockholm",
		Namn: "Oliver",
	})
	employees = append(employees, data.Employee{
		Id:   2,
		Age:  51,
		City: "Stockholm",
		Namn: "Stefam",
	})

	router := gin.Default()
	router.GET("/api/employee", apiEmployee)
	router.GET("/api/employee/:id", apiEmployeeById)
	//.POST
	router.Run(":8080")
}
