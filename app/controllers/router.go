package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/e-library/app/middleware"
)

func (h *NewHttp) Routes(r *gin.Engine) {
	r.Use(middleware.Session())
	r.Use(gin.Logger())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello"})
	})

	book := r.Group("library/book")
	{
		book.POST("/add-book", h.AddBook())
		book.GET("/get-all-books", h.GetAllBooks())
		book.PUT("/update/:book-title/:status", h.UpdateBookStatus())
		book.DELETE("/delete/:book-id", h.DeleteBook())
	}

	student := r.Group("library/student")
	{
		student.GET("/borrow/:book-title", h.BorrowBook())
		student.POST("/return-book/:student-id/:book-title", h.ReturnBook())
		student.POST("/check-in", h.CheckIn())
	}

	lend := r.Group("library/lend")
	{
		lend.GET("/get-lenders", h.GetAllBorrowedBooks())
	}

}
