package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AkshayDawkhar/Task-Tracker/db"
	"github.com/AkshayDawkhar/Task-Tracker/models"
)

func HomeH(w http.ResponseWriter, r *http.Request) {
	m := models.Massage1{Name: "akshay", Age: 20}
	b, _ := json.Marshal(m)
	w.Write(b)
}

func HomeHpost(w http.ResponseWriter, r *http.Request) {
	db.DBC.Create(&models.Massage{Name: "akshay", Age: 12, Class: 12})

}
