package models

import (
	"context"
	"fmt"
	"golang1/pkg/config"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type Employee struct {
	EmployeeName string   `json:"empName"`
	Age          int      `json:"age"`
	Reporting    string   `json:"reporting"`
	EmployeeId   string   `json:"empID"`
	Project      *Project `json:"project"`
}

type Project struct {
	ProjectName     string `json:"projName"`
	ProjectLocation string `json:"projLoc"`
	ProjectID       int    `json:"projID"`
	ProjectStatus   string `json:"projStatus"`
}

// var CNX = config.Connection()
var COLL = config.Connection().Database("youtube-employee").Collection("details")

func (e *Employee) CreateAllEmployeeDetails() *Employee {

	// coll := CNX.Database("youtube-employee").Collection("details")

	_, err := COLL.InsertOne(context.TODO(), e)

	if err != nil {
		log.Fatal(err)
	}
	return e

}

func ShowAllEmployeeDetails() []Employee {
	cursor, err := COLL.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var Employees []Employee

	if err = cursor.All(context.TODO(), &Employees); err != nil {
		log.Fatal(err)
	}
	return Employees

}

func ShowEmployeeDetails(id string) *Employee{
	var Emplyee Employee
	// database variable adding place
	cursor := COLL.FindOne(context.TODO(), bson.M{"employeeid": id}).Decode(&Emplyee)
	// cursor.Decode(&Emplyee)
	fmt.Println(cursor)
	return &Emplyee
}


func (e *Employee) UpdateEmployeeDetails (Id string) *Employee{
	var Employeess Employee 
	update := bson.M{
		"$set" : e,
	}
	_,err := COLL.UpdateOne(context.TODO(), bson.M{"employeeid" : Id}, update)
	// fmt.Println(cursor)
	if err != nil{
		panic(err)
	}
	cursor1 := COLL.FindOne(context.TODO(), bson.M{"employeeid": Id})
	cursor1.Decode(&Employeess)
	return &Employeess


}


func DeleteEmployeeDetails(Id string) []Employee{
	
	_, err := COLL.DeleteOne(context.TODO(), bson.M{"employeeid": Id})
	if err !=nil{
			panic(err)
	}
	cursor, err := COLL.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var Employees []Employee

	if err = cursor.All(context.TODO(), &Employees); err != nil {
		log.Fatal(err)
	}
	return Employees
	



}