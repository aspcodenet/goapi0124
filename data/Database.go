package data

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	DB, _ = gorm.Open(sqlite.Open("databasen.sqlite"), &gorm.Config{})
	DB.AutoMigrate(&Employee{})
	var antal int64
	DB.Model(&Employee{}).Count(&antal) // Seed SELECT COUNT(*) FROM Employee
	if antal == 0 {
		DB.Create(&Employee{Age: 50, Namn: "Stefan", City: "Test"})
		DB.Create(&Employee{Age: 14, Namn: "Oliver", City: "Test"})
		DB.Create(&Employee{Age: 20, Namn: "Josefine", City: "Test"})
	}
}
