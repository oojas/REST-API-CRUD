package DB

import (
	"CRUDAPI/Global"
	"CRUDAPI/Modals"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var urlSN = "root:admin@tcp(localhost:3306)/mydb"
var err error

func DataMigration() {
	Global.Database, err = gorm.Open(mysql.Open(urlSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Connection Failed")
	}
	Global.Database.AutoMigrate(&Modals.Employee{})
}
