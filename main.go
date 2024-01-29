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

func apiEmployeeNew(c *gin.Context) {
	// Det kommer in en ny exmployee som JSON
	// vi ska ta JSON -> employee

	// stoppa in den i slicen

	// try{
	// 	c.BindJSON(&employee)
	// 	Dataabsesdfjklasdf(&employee)
	// }
	// catch{
	// 	dfdfsfs
	// }

	// "nice att slippa exceptions -> exceptions är slow och krävande"
	var employee data.Employee
	// err := c.BindJSON(&employee)
	// if err != nil {

	// }

	if err := c.BindJSON(&employee); err != nil {
		c.AbortWithStatus(400)
		return
	}
	employee.Id = 123
	employees = append(employees, employee)
	c.IndentedJSON(http.StatusCreated, employee)
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

func apiEmployeeUpdate(c *gin.Context) {
	id := c.Param("id")

	for index, emp := range employees {
		i, _ := strconv.Atoi(id)
		if emp.Id == i {
			c.BindJSON(&employees[index])
			c.JSON(http.StatusOK, emp)
		}
	}

	c.String(http.StatusNotFound, "")
}

func main() {
	data.Init()
	// employees = append(employees, data.Employee{
	// 	Id:   1,
	// 	Age:  15,
	// 	City: "Stockholm",
	// 	Namn: "Oliver",
	// })
	// employees = append(employees, data.Employee{
	// 	Id:   2,
	// 	Age:  51,
	// 	City: "Stockholm",
	// 	Namn: "Stefam",
	// })

	router := gin.Default()
	router.GET("/api/employee", apiEmployee)
	router.GET("/api/employee/:id", apiEmployeeById)

	// NY!!!
	router.POST("/api/employee", apiEmployeeNew)
	router.PUT("/api/employee/:id", apiEmployeeUpdate)
	//.POST
	router.Run(":8080")
}
