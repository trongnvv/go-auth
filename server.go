package main

import (
	"fmt"
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

	fmt.Println("Start router ::" + PORT)
	routes.Setup(PORT)
}
