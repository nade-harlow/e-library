package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nade-harlow/e-library/app/models"
	"net/http"
	"time"
)

type NewHttp struct {
	Db    models.Db
	Route *gin.Engine
}

func New(model models.Db) *NewHttp {
	return &NewHttp{Db: model}
}

func (h *NewHttp) AddBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		book := models.Book{}
		err := c.ShouldBindJSON(&book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		book.ID = uuid.NewString()
		book.CreatedAt = time.Now().String()
		book.ModifiedAt = time.Now().String()
		fmt.Println(book.ID)
		err = h.Db.Create(book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "book added successfully"})
	}
}

func (h *NewHttp) GetAllBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := h.Db.AllBooks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"Books": books})
	}
}
