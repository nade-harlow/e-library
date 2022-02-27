package controllers

import (
	"github.com/gin-gonic/gin"
)

func (h *NewHttp) Routes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello"})
	})
	r.POST("/add", h.AddBook())
	r.GET("/get-all-books", h.GetAllBooks())
	r.POST("/check-in", h.CheckIn())
	r.POST("/borrow/:student-id/:book-title", h.BorrowBook())
	r.POST("/return-book/:student-id/:book-title", h.ReturnBook())
	r.GET("/lenders", h.GetAllBorrowedBooks())
	r.PUT("/update/:book-title/:status", h.UpdateBookStatus())
}
