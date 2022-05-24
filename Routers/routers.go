package Routers

import (
	"CRUDAPI/Handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":3000"
}

func Routers() {
	p := getPort()
	router := mux.NewRouter()
	router.HandleFunc("/createEmployees", Handlers.CreateEmployees).Methods(http.MethodPost)        // Creating employees
	router.HandleFunc("/employees", Handlers.GetEmployees).Methods(http.MethodGet)                  // Getting the employees
	router.HandleFunc("/employee/{eid}", Handlers.GetEmployeesByID).Methods(http.MethodGet)         // Getting by Id
	router.HandleFunc("/updateEmployees/{eid}", Handlers.UpdateEmployees).Methods(http.MethodPut)   // Updating the employees
	router.HandleFunc("/deleteEmployee/{eid}", Handlers.DeleteEmployees).Methods(http.MethodDelete) // Delete employee
	log.Fatal(http.ListenAndServe(p, router))
}
