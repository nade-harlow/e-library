package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/e-library/app/middleware"
)

func (h *NewHttp) Routes(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.GET("/", h.Home())
	admin := r.Group("library/admin")
	{
		admin.GET("/", h.GetAllLibraryBooks())
		admin.GET("/books", h.GetAllLibraryBooks())
		admin.GET("/add-book", h.Book())
		admin.GET("/books/history", h.GetLendingHistory())
		admin.POST("/books/history", h.GetLendingHistory())
		admin.GET("/books/:message", h.GetAllLibraryBooks())
		admin.GET("/books/delete/:book-id", h.DeleteBook())
	}

	r.GET("library/book/get-all-books/:message", h.GetAllBooks())
	r.GET("library/book/get-all-books", h.GetAllBooks())
	r.GET("/library/login", h.Login())
	r.GET("/library/logout", h.Logout())
	r.POST("library/login/auth", h.LoginAuth())
	r.GET("/library/signup", h.SignUp())
	r.POST("library/signup/auth", h.SignUpAuth())
	r.NoRoute(func(c *gin.Context) { c.HTML(404, "404.page.html", nil) })

	book := r.Group("library/book", middleware.Session())
	{
		book.POST("/add-book", h.AddBook())

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
