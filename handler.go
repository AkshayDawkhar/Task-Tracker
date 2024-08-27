package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HomeH(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1>hello</h1>"))
	// io.WriteString(w, "hiii")
	m := Massage1{"akshay", 20}
	b, _ := json.Marshal(m)
	w.Write(b)

	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Content-Typer", "application/json")
}

func HomeHpost(w http.ResponseWriter, r *http.Request) {
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
	DBC.Create(&Massage{Name: "akshay", Age: 12, class: 12})
	// a := json.Unmarshal(body,a)
	// var ms Massage
	json.NewDecoder(body).Decode(&m)
	fmt.Println(m)

}
