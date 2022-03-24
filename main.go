package main

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.com/nade-harlow/e-library/app/controllers/server"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	server.Start()
}
