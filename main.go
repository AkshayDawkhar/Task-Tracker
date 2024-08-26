package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Massage struct {
	gorm.Model

	Name string `json:"name"`
	Age  int16
	// Time int16
}
type Massage1 struct {
	Name string `json:"name"`
	Age  int16
	// Time int16
}

func homeH(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1>hello</h1>"))
	// io.WriteString(w, "hiii")
	m := Massage1{"akshay", 20}
	b, _ := json.Marshal(m)
	w.Write(b)
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Content-Typer", "application/json")
}

func homeHpost(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1>hello</h1>"))
	// io.WriteString(w, "hiii")
	fmt.Println(r)
	m := Massage1{"akshay1", 20}
	b, _ := json.Marshal(m)
	fmt.Println(b)
	w.Write(b)
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Content-Typer", "application/json")
	body := r.Body
	fmt.Println(body)
	// a := json.Unmarshal(body,a)
	// var ms Massage
	json.NewDecoder(body).Decode(&m)
	fmt.Println(m)

}
func main() {
	fmt.Println("hello")
	// r := mux.NewRouter()
	a := http.NewServeMux()
	a.HandleFunc("GET /", homeH)
	a.HandleFunc("POST /", homeHpost)
	http.ListenAndServe(":8000", a)
	// http.ListenAndServeTLS()
}
