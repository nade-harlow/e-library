package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/e-library/app/controllers"
	"github.com/nade-harlow/e-library/app/models"
	"log"
	"net/http"
	"os"
)

func Start() {
	router := gin.New()
	router.StaticFS("static", http.Dir("app/views/static"))
	router.LoadHTMLGlob("app/views/html/*")
	db := models.Init()
	dbInstance := models.NewInstance(db)
	Newhttp := controllers.New(dbInstance)
	Newhttp.Routes(router)
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port += "8080"
	}
	log.Printf("Server listening on port %s\n", port)
	err := router.Run(port)
	if err != nil {
		log.Fatalf("server failed to listen on port %s", port)
	}
}
