package main

import (
	"context"
	"fmt"
	"gostart/database"
	"gostart/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	MONGODB_URI := os.Getenv("MONGODB_URI")
	MONGODB_NAME := os.Getenv("MONGODB_NAME")

	fmt.Println("Start router ::" + PORT)

	//Setup database
	database.DB = database.Setup(MONGODB_URI, MONGODB_NAME)
	defer database.DB.Client().Disconnect(context.TODO())

	routes.Setup(PORT)
}
