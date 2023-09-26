package routes

import (
	"golang1/pkg/controller"

	"github.com/gorilla/mux"
)

var EmployeeRoutes = func(r *mux.Router){
	
	r.HandleFunc("/employees", controller.ShowAllEmployeeDetails).Methods("GET")
	r.HandleFunc("/employee/{id}", controller.ShowEmployeeDetails).Methods("GET")
	r.HandleFunc("/employee", controller.CreateAllEmployeeDetails).Methods("POST")
	r.HandleFunc("/employee/{id}", controller.UpdateEmployeeDetails).Methods("PUT")
	r.HandleFunc("/employee/{id}", controller.DeleteEmployeeDetails).Methods("DELETE")
}

