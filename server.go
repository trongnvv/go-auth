package main

import (
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
	URI_MONGODB := os.Getenv("URI_MONGODB")

	fmt.Println("Start router ::" + PORT)

	db := database.Setup(URI_MONGODB, "trongnv")
	routes.Setup(PORT, db)
}
