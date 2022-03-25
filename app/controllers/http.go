package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nade-harlow/e-library/app/helper"
	"github.com/nade-harlow/e-library/app/models"
	"log"
	"net/http"
	"strings"
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
		c.HTML(200, "login.html", nil)
	}
}

func (h NewHttp) Book() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "admin.add.books.html", nil)
	}
}

func (h NewHttp) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("session", "", -1, "", "", true, true)
		c.Redirect(http.StatusFound, "/library/book/get-all-books")
	}
}

func (h NewHttp) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	}
}

func (h NewHttp) LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.PostForm("user_name")
		password := c.PostForm("password")
		user, err := h.Db.StudentLogin(userName, password)
		if err != nil {
			log.Println(err.Error())
			return
		}
		if user.FirstName != "" {
			log.Println(user.ID)
			c.SetCookie("session", user.ID, 3600, "/", "", true, true)
			c.Redirect(http.StatusFound, "/library/lend/get-lenders")
			return
		}
		//TODO: handle constraint
		c.Redirect(http.StatusFound, "/library/signup")
	}
}

func (h NewHttp) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "signup.html", nil)
	}
}

func (h *NewHttp) SignUpAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var student models.Student
		student.ID = uuid.NewString()
		student.FirstName = strings.ToLower(c.PostForm("first_name"))
		student.LastName = strings.ToLower(c.PostForm("last_name"))
		student.UserName = strings.ToLower(c.PostForm("user_name"))
		student.Password = c.PostForm("password")
		confirmPassword := c.PostForm("c_password")
		if student.Password != confirmPassword {
			// TODO: handle constraints
			log.Println("password doesn't match")
			return
		}
		if err := h.Db.StudentSignUp(student); err != nil {
			// TODO: handle constraints
			log.Println("signup error", err.Error())
			return
		}
		c.Redirect(http.StatusFound, "/library/login")
	}
}

func (h *NewHttp) AddBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		book := models.Book{}
		//title := c.PostFormArray("title")
		//author := c.PostFormArray("author")
		//url := c.PostFormArray("url")
		//stock := c.PostFormArray("stock")
		//log.Println("title", title)
		//log.Println("author", author)
		//log.Println("url", url)
		//log.Println("stock", stock)
		book.Title = c.PostForm("title")
		book.Author = c.PostForm("author")
		book.Url = c.PostForm("url")
		err := h.Db.AddBook(book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Redirect(http.StatusFound, "/library/admin/books")
	}
}

func (h *NewHttp) GetAllBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		message := c.Param("message")
		books, err := h.Db.GetAllBooks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.HTML(200, "library.books.html", gin.H{"Books": books, "Message": message})
	}
}

func (h *NewHttp) CheckIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		student := models.Student{}
		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")
		userName := c.PostForm("user_name")
		password := c.PostForm("password")
		student.FirstName = firstName
		student.LastName = lastName
		student.UserName = userName
		student.Password = password
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
			c.SetCookie("session", student.ID, 3600, "/", "", true, true)
			c.Redirect(http.StatusFound, "/library/book/get-all-books")
			return
		}
		c.SetCookie("session", data.ID, 3600, "/", "", true, true)
		c.Redirect(http.StatusFound, "/library/book/get-all-books")
	}
}

func (h NewHttp) BorrowBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		//studentID:= "1"
		student, exist := c.Get("student")
		studentID := student.(string)
		if !exist || studentID == "" {
			c.Redirect(http.StatusFound, "/library/student/check-in")
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
			message := "you've already borrowed this book"
			c.Redirect(http.StatusFound, fmt.Sprintf("/library/book/get-all-books/%s", message))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		message := fmt.Sprintf(`you just borrowed %s by %s`, strings.ToTitle(book.Title), strings.ToTitle(book.Author))
		c.Redirect(http.StatusFound, fmt.Sprintf("/library/book/get-all-books/%s", message))
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
		c.Redirect(http.StatusFound, "/library/lend/get-lenders")
		c.JSON(http.StatusOK, gin.H{"response": fmt.Sprintf(`Thank you for returning '%s'`, book.Title)})
	}
}

func (h NewHttp) GetAllBorrowedBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		//studentID:= "1"
		student, exist := c.Get("student")
		studentID := student.(string)

		if !exist || studentID == "" {
			c.Redirect(http.StatusFound, "/library/student/check-in")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "sorry, you have to check in first"})
			return
		}
		books, err := h.Db.GetBorrowedBooks(studentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.HTML(200, "return.book.portal.html", gin.H{"Borrowed": books})
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
		message := "Book Removed From Library Successfully"
		c.Redirect(http.StatusFound, fmt.Sprintf("/library/admin/books/%s", message))
	}
}

func (h NewHttp) Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.PostForm("book-title")
		book, err := h.Db.GetBook(title)
		if err != nil {
			log.Println(err)
		}
		c.HTML(200, "search.result.html", book)
	}
}

func (h NewHttp) GetLendingHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var returned bool
		filter := strings.ToLower(c.PostForm("filter"))
		if filter == "returned" {
			returned = true
		}
		allLending, err := h.Db.GetAllLending(returned)
		if err != nil {
			return
		}
		c.HTML(200, "admin.books.history.html", gin.H{"History": allLending})
	}
}

// GetAllLibraryBooks Admin portal
func (h *NewHttp) GetAllLibraryBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		message := c.Param("message")
		books, err := h.Db.GetAllBooks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.HTML(200, "admin.all.books.html", gin.H{"Books": books, "Message": message})
	}
}
