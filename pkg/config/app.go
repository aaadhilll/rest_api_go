package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() *mongo.Client{
	if err := godotenv.Load(); err != nil{
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")

	if uri == ""{
		log.Fatal("Sorry uri is blanck")
	}
	// set cline options
	clinetOprions := options.Client().ApplyURI(uri)
	// creating mongo connections
	clinet, err := mongo.Connect(context.TODO(), clinetOprions)

	if err != nil{
		log.Fatal(err)
	}

	// check connecton
	
	err = clinet.Ping(context.TODO(), nil)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("connected to mingo db")

	return clinet
		

}