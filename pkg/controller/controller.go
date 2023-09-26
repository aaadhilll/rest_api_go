package controller

import (
	"encoding/json"
	"fmt"
	"golang1/pkg/models"
	"golang1/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func ShowAllEmployeeDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("within functions")

	// fmt.Fprintf(w, "hello")
	showEmployee := models.ShowAllEmployeeDetails()
	res,_ := json.Marshal(showEmployee)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ShowEmployeeDetails(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, "hello")
	vars := mux.Vars(r)
	empID := vars["id"]
	empDetails := models.ShowEmployeeDetails(empID)
	res,_ := json.Marshal(empDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateAllEmployeeDetails(w http.ResponseWriter, r *http.Request) {
	CreateEmployee := &models.Employee{}
	// fmt.Println("bodeeeeee")
	// fmt.Println(r.)
	utils.ParseBody(r, CreateEmployee)
	employee := CreateEmployee.CreateAllEmployeeDetails()
	res,_ := json.Marshal(employee)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// fmt.Fprintf(w, "hello")
}
func UpdateEmployeeDetails(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "hello")
	vars := mux.Vars(r)
	empID := vars["id"]
	UpdateEmployee := &models.Employee{}
	// fmt.Println("bodeeeeee")
	// fmt.Println(r.)
	utils.ParseBody(r, UpdateEmployee)
	empDetailsToUpdate := UpdateEmployee.UpdateEmployeeDetails(empID)
	res,_ := json.Marshal(empDetailsToUpdate)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteEmployeeDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID := vars["id"]
	empDetails := models.DeleteEmployeeDetails(empID)
	res,_ := json.Marshal(empDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

	
}