package db

import (
	"github.com/AkshayDawkhar/Task-Tracker/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBC *gorm.DB

func ConnectDB() {
	var err error
	DBC, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	DBC.AutoMigrate(&models.Massage{})

	DBC.Create(&models.Massage{Name: "akshay", Age: 12, Class: 12})
}
