package main

import (
	"CRUDAPI/DB"
	"CRUDAPI/Routers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	DB.DataMigration()
	Routers.Routers()
}
