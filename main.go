package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
	"github.com/nade-harlow/e-library/app/controllers"
	"github.com/nade-harlow/e-library/app/models"
	"log"
	"os"
)

func main() {
	router := gin.Default()
	db := models.Init()
	//di:= models.DbInstance{Postgres: db}
	dbIns := models.NewInstance(db)
	//c := controllers.NewHttp{&di}
	Nh := controllers.New(dbIns)
	Nh.Routes(router)
	port := os.Getenv("DB_HOST")
	log.Printf("Server listening on port %s\n", port)
	err := router.Run()
	if err != nil {
		log.Fatalf("server failed to listen on port %s", port)
	}
}
