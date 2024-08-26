package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello")
	// r := mux.NewRouter()
	a := http.NewServeMux()
	a.HandleFunc("GET /", HomeH)
	a.HandleFunc("POST /", HomeHpost)
	http.ListenAndServe(":8000", a)
	// http.ListenAndServeTLS()
}
