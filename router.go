package main

import "net/http"

func SetupRoute(router *http.ServeMux) {
	router.HandleFunc("GET /", HomeH)
	router.HandleFunc("POST /", HomeHpost)
}
