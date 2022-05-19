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
	return ":8000"
}

func Routers() {
	p := getPort()
	router := mux.NewRouter()
	router.HandleFunc("/createEmployees", Handlers.CreateEmployees).Methods(http.MethodPost)
	router.HandleFunc("/employees", Handlers.GetEmployees).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(p, router))
}
