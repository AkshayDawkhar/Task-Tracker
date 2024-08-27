package main

import (
	"fmt"
	"net/http"

	"github.com/AkshayDawkhar/Task-Tracker/models"
	"github.com/AkshayDawkhar/Task-Tracker/routers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBC *gorm.DB

func main() {
	fmt.Println("hello")
	// r := mux.NewRouter()
	DBC, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	DBC.AutoMigrate(&models.Massage{})

	DBC.Create(&models.Massage{Name: "akshay", Age: 12, Class: 12})

	router := http.NewServeMux()
	routers.SetupRoute(router)
	http.ListenAndServe(":8000", router)
	// http.ListenAndServeTLS()
}
