package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func (h *NewHttp) Routes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		log.Println(os.Getenv("DB_PORT"))
		c.JSON(200, gin.H{"message": "hello"})
	})
	r.POST("/add", h.AddBook())
	r.GET("/get-all-books", h.GetAllBooks())
}
