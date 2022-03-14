package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
	"github.com/nade-harlow/e-library/app/controllers"
	"github.com/nade-harlow/e-library/app/models"
	"log"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	router.StaticFS("static", http.Dir("app/views/static"))
	router.LoadHTMLGlob("app/views/html/*")
	db := models.Init()
	dbInstance := models.NewInstance(db)
	Newhttp := controllers.New(dbInstance)
	Newhttp.Routes(router)
	port := os.Getenv("DB_HOST")
	log.Printf("Server listening on port %s\n", port)
	err := router.Run()
	if err != nil {
		log.Fatalf("server failed to listen on port %s", port)
	}
}
