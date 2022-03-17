package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/e-library/app/middleware"
)

func (h *NewHttp) Routes(r *gin.Engine) {
	//r.Use(middleware.Session())
	r.Use(gin.Recovery())
	r.GET("/", h.Home())
	book := r.Group("library/book", middleware.Session())
	{
		book.POST("/add-book", h.AddBook())
		book.GET("/get-all-books/:message", h.GetAllBooks())
		book.GET("/get-all-books", h.GetAllBooks())
		book.POST("/search", h.Search())
		book.GET("/search/result", h.Search())
		book.PUT("/update/:book-title/:status", h.UpdateBookStatus())
		book.DELETE("/delete/:book-id", h.DeleteBook())
	}

	student := r.Group("library/student", middleware.Session())
	{
		student.GET("/borrow/:book-title", h.BorrowBook())
		student.GET("/return-book/:student-id/:book-title", h.ReturnBook())
		student.GET("/check-in", h.Home())
		student.POST("/check-in", h.CheckIn())
	}

	lend := r.Group("library/lend", middleware.Session())
	{
		lend.GET("/get-lenders", h.GetAllBorrowedBooks())
	}
}
