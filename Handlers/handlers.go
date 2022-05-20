package Handlers

import (
	"CRUDAPI/Global"
	"CRUDAPI/Modals"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var emp Modals.Employee
	json.NewDecoder(r.Body).Decode(&emp)
	Global.Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)
}
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var employees []Modals.Employee
	Global.Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)
}
func GetEmployeesByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var employees Modals.Employee
	Global.Database.First(&employees, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(employees)
}
