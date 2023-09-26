package main

import (
	"fmt"
	"golang1/pkg/routes"
	"log"
	"net/http"
	"github.com/gorilla/mux"

)

func main(){
	 fmt.Println("1")
	 
	 r := mux.NewRouter()
	 fmt.Println("2")
	 routes.EmployeeRoutes(r)
	 fmt.Println("3")
	 http.Handle("/", r)
	 fmt.Println("4")
	
	 fmt.Println("server running on: http://localhost:5003")
	 log.Fatal(http.ListenAndServe(":5003", r))
	 fmt.Println("5")
}