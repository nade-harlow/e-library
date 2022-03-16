package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	db "github.com/nade-harlow/e-library/app/mocks"
	"github.com/nade-harlow/e-library/app/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Loader(mdb *db.MockDb) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("../views/html/*")
	newhttp := &NewHttp{
		Db:    mdb,
		Route: router,
	}
	newhttp.Routes(router)
	return router
}

func TestNewHttp_AddBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mdb := db.NewMockDb(ctrl)
	router := Loader(mdb)

	book := models.Book{
		ID:         "1",
		Title:      "women of owo",
		Author:     "kwame",
		Available:  true,
		CreatedAt:  "1pm",
		ModifiedAt: "1pm",
	}
	mdb.EXPECT().AddBook(book).Return(errors.New("can't insert to db"))
	mdb.EXPECT().AddBook(book).Return(nil)
	body, err := json.Marshal(&book)
	if err != nil {
		t.Fail()
		return
	}
	t.Run("testing error", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodPost, "/library/book/add-book", strings.NewReader(string(body)))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		res, err := json.Marshal(gin.H{"error": "can't insert to db"})
		if err != nil {
			t.Fatal(err)
		}
		if response.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, response.Code)
		}
		if string(response.Body.Bytes()) != string(res) {
			t.Errorf("Expected %s, got %s", `{"error": "can't insert to db"}`, string(response.Body.Bytes()))
		}
	})

	t.Run("testing no error", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodPost, "/library/book/add-book", strings.NewReader(string(body)))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		res, err := json.Marshal(gin.H{"message": "book added successfully"})
		if err != nil {
			t.Fatal(err)
		}
		if response.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
		if string(response.Body.Bytes()) != string(res) {
			t.Errorf("Expected %s, got %s", `{"message":"book added successfully"}`, string(response.Body.Bytes()))
		}
	})

}

func TestNewHttp_GetAllBooks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mdb := db.NewMockDb(ctrl)
	router := Loader(mdb)
	book := []models.Book{
		{
			ID:         "1",
			Title:      "women of owo",
			Author:     "kwame",
			Available:  true,
			CreatedAt:  "1pm",
			ModifiedAt: "1pm",
		},
		{
			ID:         "2",
			Title:      "arms and the man",
			Author:     "bernard shaw",
			Available:  true,
			CreatedAt:  "3pm",
			ModifiedAt: "3pm",
		},
	}
	mdb.EXPECT().GetAllBooks().Return(nil, errors.New("some error getting books"))
	mdb.EXPECT().GetAllBooks().Return(book, nil)
	t.Run("testing error getting books", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/library/book/get-all-books", nil)
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "text/html")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, response.Code)
		}
	})

	t.Run("testing getting all books success", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/library/book/get-all-books", nil)
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "text/html")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res := `<p class="card-text">women of owo</p>`
		if response.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
		if !strings.Contains(response.Body.String(), res) {
			t.Errorf("Expected %s, got %s", `<p class="card-text">women of owo</p>`, response.Body.String())
		}
	})

}

func TestNewHttp_CheckIn(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mdb := db.NewMockDb(ctrl)
	router := Loader(mdb)
	mdb.EXPECT().GetStudentByName("jim", "morrison").Return(models.Student{}, nil)
	mdb.EXPECT().StudentCheckIn(gomock.Any()).Return(errors.New("error checking student in"))
	//mdb.EXPECT().StudentCheckIn(gomock.Any()).Return(nil)

	form := url.Values{}
	form.Set("first_name", "jim")
	form.Set("last_name", "morrison")

	t.Run("testing error checking student in", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodPost, "/library/student/check-in", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, response.Code)
		}
	})

	//t.Run("testing checking student in success", func(t *testing.T) {
	//	request, err := http.NewRequest(http.MethodPost, "/library/student/check-in", strings.NewReader(form.Encode()))
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//	response := httptest.NewRecorder()
	//	router.ServeHTTP(response, request)
	//
	//	if response.Code != http.StatusFound {
	//		t.Errorf("Expected status code %d, got %d", http.StatusFound, response.Code)
	//	}
	//	//if string(response.Body.Bytes()) != string(res) {
	//	//	t.Errorf("Expected %s, got %s", `{"response": "Successfully checked in. Happy reading"}`, string(response.Body.Bytes()))
	//	//}
	//})

}

func TestNewHttp_BorrowBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mdb := db.NewMockDb(ctrl)
	router := gin.Default()
	newhttp := &NewHttp{
		Db:    mdb,
		Route: router,
	}
	newhttp.Routes(router)

	book := models.Book{
		ID:         "1",
		Title:      "women of owo",
		Author:     "kwame",
		Available:  true,
		CreatedAt:  "1pm",
		ModifiedAt: "1pm",
	}
	mdb.EXPECT().GetBookByTitle(gomock.Any()).Return(book, nil)
	mdb.EXPECT().BorrowBook(book.ID, gomock.Any()).Return(nil)

	request, err := http.NewRequest(http.MethodGet, "/library/student/borrow/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	res, err := json.Marshal(gin.H{"response": fmt.Sprintf(`you just borrowed '%s' by '%s'`, book.Title, book.Author)})
	if err != nil {
		t.Fatal(err)
	}
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
	}
	if string(response.Body.Bytes()) != string(res) {
		t.Errorf("Expected %s, got %s", `{"response": fmt.Sprintf("you just borrowed '%s' by '%s'", book.Title, book.Author)}`, string(response.Body.Bytes()))
	}
}

func TestNewHttp_ReturnBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mdb := db.NewMockDb(ctrl)
	router := gin.Default()
	newhttp := &NewHttp{
		Db:    mdb,
		Route: router,
	}
	newhttp.Routes(router)

	studentID := "1"
	book := models.Book{
		ID:         "1",
		Title:      "women of owo",
		Author:     "kwame",
		Available:  true,
		CreatedAt:  "1pm",
		ModifiedAt: "1pm",
	}
	mdb.EXPECT().GetBookByTitle(gomock.Any()).Return(book, nil)
	mdb.EXPECT().ReturnBook(studentID, book.ID).Return(nil)

	body, err := json.Marshal(&book)
	if err != nil {
		t.Fail()
	}
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/library/student/return-book/%s/%s", studentID, book.Title), strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	res, err := json.Marshal(gin.H{"response": fmt.Sprintf(`Thank you for returning '%s'`, book.Title)})
	if err != nil {
		t.Fatal(err)
	}
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
	}
	if string(response.Body.Bytes()) != string(res) {
		t.Errorf("Expected %s, got %s", `{"response": fmt.Sprintf("Thank you for returning '%s'", book.Title)}`, string(response.Body.Bytes()))
	}
}

func TestNewHttp_GetAllBorrowedBooks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mdb := db.NewMockDb(ctrl)
	router := gin.Default()
	newhttp := &NewHttp{
		Db:    mdb,
		Route: router,
	}
	newhttp.Routes(router)

	lendings := []models.BorrowedBook{
		{
			ID:         "1",
			StudentID:  "1",
			BookID:     "1",
			Returned:   false,
			CreatedAt:  "1pm",
			ModifiedAt: "1pm",
		},
	}
	mdb.EXPECT().GetAllLending().Return(lendings, nil)

	request, err := http.NewRequest(http.MethodGet, "/library/lend/get-lenders", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	res, err := json.Marshal(gin.H{"Lenders": lendings})
	if err != nil {
		t.Fatal(err)
	}
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
	}
	if string(response.Body.Bytes()) != string(res) {
		t.Errorf("Expected %v, got %s", gin.H{"Lenders": lendings}, string(response.Body.Bytes()))
	}
}

func TestNewHttp_UpdateBookStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mdb := db.NewMockDb(ctrl)
	router := gin.Default()
	newhttp := &NewHttp{
		Db:    mdb,
		Route: router,
	}
	newhttp.Routes(router)
	status := "available" // available = true || not available = false
	book := models.Book{
		ID:         "1",
		Title:      "women of owo",
		Author:     "kwame",
		Available:  true,
		CreatedAt:  "1pm",
		ModifiedAt: "1pm",
	}
	mdb.EXPECT().GetBook(book.Title).Return(book, nil)
	mdb.EXPECT().UpdateBookStatus(true, book.ID).Return(nil)

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/library/book/update/%s/%s", book.Title, status), nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	res, err := json.Marshal(gin.H{"response": "book status updated successfully"})
	if err != nil {
		t.Fatal(err)
	}
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
	}
	if string(response.Body.Bytes()) != string(res) {
		t.Errorf("Expected %v, got %s", gin.H{"response": "book status updated successfully"}, string(response.Body.Bytes()))
	}
}

func TestNewHttp_DeleteBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mdb := db.NewMockDb(ctrl)
	router := gin.Default()
	newhttp := &NewHttp{
		Db:    mdb,
		Route: router,
	}
	newhttp.Routes(router)

	bookID := "1"
	mdb.EXPECT().DeleteBookById(bookID).Return(errors.New("error deleting book"))
	mdb.EXPECT().DeleteBookById(bookID).Return(nil)

	t.Run("testing error deleting book", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/library/book/delete/%s", bookID), nil)
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, response.Code)
		}
	})

	t.Run("testing delete book success", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/library/book/delete/%s", bookID), nil)
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		res, err := json.Marshal(gin.H{"response": "book has been deleted successfully"})
		if err != nil {
			t.Fatal(err)
		}
		if response.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
		if string(response.Body.Bytes()) != string(res) {
			t.Errorf("Expected %v, got %s", gin.H{"response": "book has been deleted successfully"}, string(response.Body.Bytes()))
		}
	})

}
