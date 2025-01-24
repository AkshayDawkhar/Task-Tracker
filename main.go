package main

import (
	"fmt"
	"net/http"

	"github.com/AkshayDawkhar/Task-Tracker/routers"
	"gorm.io/gorm"
)

var DBC *gorm.DB

func main() {
	fmt.Println("starting server")
	router := http.NewServeMux()
	routers.SetupRoute(router)
	http.ListenAndServe(":8000", router)
	// http.ListenAndServeTLS()
}
