package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AkshayDawkhar/Task-Tracker/models"
)

func HomeH(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1>hello</h1>"))
	// io.WriteString(w, "hiii")
	m := models.Massage1{Name: "akshay", Age: 20}
	b, _ := json.Marshal(m)
	w.Write(b)

	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Content-Typer", "application/json")
}

func HomeHpost(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1>hello</h1>"))
	// io.WriteString(w, "hiii")
	fmt.Println(r)
	m := models.Massage1{Name: "akshay1", Age: 20}
	b, _ := json.Marshal(m)
	fmt.Println(b)
	w.Write(b)
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Content-Typer", "application/json")
	body := r.Body
	fmt.Println(body)
	// main.DBC.Create(&Massage{Name: "akshay", Age: 12, class: 12})
	// a := json.Unmarshal(body,a)
	// var ms Massage
	json.NewDecoder(body).Decode(&m)
	fmt.Println(m)

}
