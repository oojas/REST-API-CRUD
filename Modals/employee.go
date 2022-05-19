package Modals

import "gorm.io/gorm"

type Employee struct {
	gorm.Model        // Helps in generating auto Id of the record as well as keeps the track of the time at which the record was uploaded /deleted.
	EmpName    string `json:"emp_name"`
	EmpSalary  int    `json:"emp_salary"`
	Email      string `json:"email"`
}
