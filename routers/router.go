package routers

import (
	"net/http"

	"github.com/AkshayDawkhar/Task-Tracker/handlers"
)

func SetupRoute(router *http.ServeMux) {
	router.HandleFunc("GET /", handlers.HomeH)
	router.HandleFunc("POST /", handlers.HomeHpost)
}
