package controllers

import (
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

func (h *NewHttp) CheckIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		student := models.Student{}
		err := c.ShouldBindJSON(&student)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		student.ID = uuid.NewString()
		err = h.Db.StudentCheckIn(student)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"response": "Successfully checked in. Happy reading"})
		c.Set("student", student.ID)
	}
}

func (h NewHttp) BorrowBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		student, exist := c.Get("student")
		if !exist {
			c.JSON(http.StatusNotFound, gin.H{"error": "sorry, you have to check in first"})
		}
		studentID := student.(*models.Student).ID
		book, err := h.Db.GetBookByTitle("Things fall apart")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = h.Db.BorrowBook(book.ID, studentID)
		if err != nil {
			return
		}
	}
}
