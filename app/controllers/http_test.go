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
	"strings"
	"testing"
)

func TestNewHttp_AddBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mdb := db.NewMockDb(ctrl)
	router := gin.Default()
	newhttp := &NewHttp{
		Db:    mdb,
		Route: router,
	}
	book := models.Book{
		ID:         "1",
		Title:      "women of owo",
		Author:     "kwame",
		Available:  true,
		CreatedAt:  "1pm",
		ModifiedAt: "1pm",
	}
	newhttp.Routes(router)
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
	router := gin.Default()
	newhttp := &NewHttp{
		Db:    mdb,
		Route: router,
	}
	newhttp.Routes(router)
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
	mdb.EXPECT().GetAllBooks().Return(book, nil)
	request, err := http.NewRequest(http.MethodGet, "/library/book/get-all-books", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	res, err := json.Marshal(gin.H{"Books": book})
	if err != nil {
		t.Fatal(err)
	}
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
	}
	if string(response.Body.Bytes()) != string(res) {
		t.Errorf("Expected %s, got %s", `{"message":"book added successfully"}`, string(response.Body.Bytes()))
	}
}

func TestNewHttp_CheckIn(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mdb := db.NewMockDb(ctrl)
	router := gin.Default()
	newhttp := &NewHttp{
		Db:    mdb,
		Route: router,
	}
	newhttp.Routes(router)
	student := models.Student{
		ID:         "1",
		FirstName:  "Jim",
		LastName:   "Morrison",
		CreatedAt:  "1pm",
		ModifiedAt: "1pm",
	}

	mdb.EXPECT().StudentCheckIn(gomock.Any()).Return(nil)
	body, err := json.Marshal(&student)
	if err != nil {
		t.Fail()
	}
	request, err := http.NewRequest(http.MethodPost, "/library/student/check-in", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	res, err := json.Marshal(gin.H{"response": "Successfully checked in. Happy reading"})
	if err != nil {
		t.Fatal(err)
	}
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
	}
	if string(response.Body.Bytes()) != string(res) {
		t.Errorf("Expected %s, got %s", `{"response": "Successfully checked in. Happy reading"}`, string(response.Body.Bytes()))
	}
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
	mdb.EXPECT().BorrowBook(book.ID, studentID).Return(nil)

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
