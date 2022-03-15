package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nade-harlow/e-library/app/helper"
	"github.com/nade-harlow/e-library/app/models"
	"log"
	"net/http"
)

type NewHttp struct {
	Db    models.Db
	Route *gin.Engine
}

func New(model models.Db) *NewHttp {
	return &NewHttp{Db: model}
}

func (h NewHttp) Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "student.checkin.html", nil)
	}
}

func (h *NewHttp) AddBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		book := models.Book{}
		err := c.ShouldBindJSON(&book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = h.Db.AddBook(book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "book added successfully"})
	}
}

func (h *NewHttp) GetAllBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := h.Db.GetAllBooks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.HTML(200, "library.books.html", gin.H{"Books": books})
		//c.JSON(200, gin.H{"Books": books})
	}
}

func (h *NewHttp) CheckIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		student := models.Student{}
		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")
		student.FirstName = firstName
		student.LastName = lastName
		student.ID = uuid.NewString()
		data, err := h.Db.GetStudentByName(firstName, lastName)
		if err != nil {
			log.Println(err.Error())
		}
		if data.ID == "" {
			err = h.Db.StudentCheckIn(student)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			helper.SaveSession(student.ID)
			c.Redirect(302, "/library/book/get-all-books")
			return
		}
		helper.SaveSession(data.ID)
		c.Redirect(302, "/library/book/get-all-books")
	}
}

func (h NewHttp) BorrowBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		student, exist := c.Get("student")
		studentID := student.(string)
		if !exist || studentID == "" {
			c.Redirect(302, "/library/student/check-in")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "sorry, you have to check in first"})
			return
		}
		bookTitle := c.Param("book-title")
		book, err := h.Db.GetBookByTitle(bookTitle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = h.Db.BorrowBook(book.ID, studentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Redirect(302, "/library/book/get-all-books")
		c.JSON(http.StatusOK, gin.H{"response": fmt.Sprintf(`you just borrowed '%s' by '%s'`, book.Title, book.Author)})
	}
}

func (h NewHttp) ReturnBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookTitle := c.Param("book-title")
		studentID := c.Param("student-id")
		book, err := h.Db.GetBookByTitle(bookTitle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = h.Db.ReturnBook(studentID, book.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Redirect(302, "/library/lend/get-lenders")
		c.JSON(http.StatusOK, gin.H{"response": fmt.Sprintf(`Thank you for returning '%s'`, book.Title)})
	}
}

func (h NewHttp) GetAllBorrowedBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		student, exist := c.Get("student")
		studentID := student.(string)
		if !exist || studentID == "" {
			c.Redirect(302, "/library/student/check-in")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "sorry, you have to check in first"})
			return
		}
		books, err := h.Db.GetBorrowedBooks(studentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.HTML(200, "return.book.portal.html", gin.H{"Borrowed": books})
		//c.JSON(http.StatusOK, gin.H{"Lenders": books})
	}
}

func (h NewHttp) UpdateBookStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Param("status")
		bookTitle := c.Param("book-title")
		bookStatus, err := helper.CheckBookStatus(status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		book, er := h.Db.GetBook(bookTitle)
		if er != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
			return
		}
		err = h.Db.UpdateBookStatus(bookStatus, book.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": "book status updated successfully"})
	}
}

func (h NewHttp) DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("book-id")
		err := h.Db.DeleteBookById(bookID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": "book has been deleted successfully"})
	}
}
