package main

import (
	"fmt"
	"net/http"

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
	DBC.AutoMigrate(&Massage{})

	DBC.Create(&Massage{Name: "akshay", Age: 12, class: 12})

	router := http.NewServeMux()
	SetupRoute(router)
	http.ListenAndServe(":8000", router)
	// http.ListenAndServeTLS()
}
