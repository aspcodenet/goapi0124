package main

// 1. När vi kör lokalt utvecklar så SqLite
//						en setting är ju filnamnet
// 2. kör i Kubernetes (prod) så MySQL -
//				servernnamnet, port, databasen, anvnamn, lösenord
//                   kubernetes SECRETS !!!

// möjlighet ar override:a med environmentvariabler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"systementor.se/goapi0124/data"
)

func apiEmployee(c *gin.Context) {
	var employees []data.Employee
	data.DB.Find(&employees)

	c.IndentedJSON(http.StatusOK, employees)
}

func apiEmployeeNew(c *gin.Context) {
	var employee data.Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	employee.Id = 0
	err := data.DB.Create(&employee).Error
	if err != nil {

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, employee)
	}
}

func apiEmployeeById(c *gin.Context) {
	id := c.Param("id")
	var employee data.Employee
	err := data.DB.First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		c.IndentedJSON(http.StatusOK, employee)
	}
}

func apiEmployeeUpdate(c *gin.Context) {
	id := c.Param("id")
	var employee data.Employee
	err := data.DB.First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		if err := c.BindJSON(&employee); err != nil {
			return
		}
		employee.Id, _ = strconv.Atoi(id)
		data.DB.Save(&employee)
		c.IndentedJSON(http.StatusOK, employee)
	}
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
