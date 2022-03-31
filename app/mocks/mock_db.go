// Code generated by MockGen. DO NOT EDIT.
// Source: app/models/db.go

// Package db is a generated GoMock package.
package db

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/nade-harlow/e-library/app/models"
)

// MockDb is a mock of Db interface.
type MockDb struct {
	ctrl     *gomock.Controller
	recorder *MockDbMockRecorder
}

// MockDbMockRecorder is the mock recorder for MockDb.
type MockDbMockRecorder struct {
	mock *MockDb
}

// NewMockDb creates a new mock instance.
func NewMockDb(ctrl *gomock.Controller) *MockDb {
	mock := &MockDb{ctrl: ctrl}
	mock.recorder = &MockDbMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDb) EXPECT() *MockDbMockRecorder {
	return m.recorder
}

// AddBook mocks base method.
func (m *MockDb) AddBook(book models.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBook", book)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBook indicates an expected call of AddBook.
func (mr *MockDbMockRecorder) AddBook(book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBook", reflect.TypeOf((*MockDb)(nil).AddBook), book)
}

// BorrowBook mocks base method.
func (m *MockDb) BorrowBook(bookId, studentId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BorrowBook", bookId, studentId)
	ret0, _ := ret[0].(error)
	return ret0
}

// BorrowBook indicates an expected call of BorrowBook.
func (mr *MockDbMockRecorder) BorrowBook(bookId, studentId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BorrowBook", reflect.TypeOf((*MockDb)(nil).BorrowBook), bookId, studentId)
}

// CheckBookAvailability mocks base method.
func (m *MockDb) CheckBookAvailability(title string) (bool, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBookAvailability", title)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// CheckBookAvailability indicates an expected call of CheckBookAvailability.
func (mr *MockDbMockRecorder) CheckBookAvailability(title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBookAvailability", reflect.TypeOf((*MockDb)(nil).CheckBookAvailability), title)
}

// CheckIfBorrowed mocks base method.
func (m *MockDb) CheckIfBorrowed(studentID, bookID string) (models.BorrowedBook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIfBorrowed", studentID, bookID)
	ret0, _ := ret[0].(models.BorrowedBook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIfBorrowed indicates an expected call of CheckIfBorrowed.
func (mr *MockDbMockRecorder) CheckIfBorrowed(studentID, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfBorrowed", reflect.TypeOf((*MockDb)(nil).CheckIfBorrowed), studentID, bookID)
}

// CheckLendStatus mocks base method.
func (m *MockDb) CheckLendStatus(studentId, bookId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckLendStatus", studentId, bookId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckLendStatus indicates an expected call of CheckLendStatus.
func (mr *MockDbMockRecorder) CheckLendStatus(studentId, bookId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckLendStatus", reflect.TypeOf((*MockDb)(nil).CheckLendStatus), studentId, bookId)
}

// CheckStockCount mocks base method.
func (m *MockDb) CheckStockCount(ID string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckStockCount", ID)
	ret0, _ := ret[0].(int)
	return ret0
}

// CheckStockCount indicates an expected call of CheckStockCount.
func (mr *MockDbMockRecorder) CheckStockCount(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckStockCount", reflect.TypeOf((*MockDb)(nil).CheckStockCount), ID)
}

// DeleteBookById mocks base method.
func (m *MockDb) DeleteBookById(bookID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBookById", bookID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBookById indicates an expected call of DeleteBookById.
func (mr *MockDbMockRecorder) DeleteBookById(bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBookById", reflect.TypeOf((*MockDb)(nil).DeleteBookById), bookID)
}

// DeleteBookByTitle mocks base method.
func (m *MockDb) DeleteBookByTitle(bookTitle string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBookByTitle", bookTitle)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBookByTitle indicates an expected call of DeleteBookByTitle.
func (mr *MockDbMockRecorder) DeleteBookByTitle(bookTitle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBookByTitle", reflect.TypeOf((*MockDb)(nil).DeleteBookByTitle), bookTitle)
}

// GetAllBooks mocks base method.
func (m *MockDb) GetAllBooks() ([]models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBooks")
	ret0, _ := ret[0].([]models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBooks indicates an expected call of GetAllBooks.
func (mr *MockDbMockRecorder) GetAllBooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBooks", reflect.TypeOf((*MockDb)(nil).GetAllBooks))
}

// GetAllLending mocks base method.
func (m *MockDb) GetAllLending(returned bool) ([]map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLending", returned)
	ret0, _ := ret[0].([]map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLending indicates an expected call of GetAllLending.
func (mr *MockDbMockRecorder) GetAllLending(returned interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLending", reflect.TypeOf((*MockDb)(nil).GetAllLending), returned)
}

// GetBook mocks base method.
func (m *MockDb) GetBook(title string) (models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBook", title)
	ret0, _ := ret[0].(models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBook indicates an expected call of GetBook.
func (mr *MockDbMockRecorder) GetBook(title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBook", reflect.TypeOf((*MockDb)(nil).GetBook), title)
}

// GetBookById mocks base method.
func (m *MockDb) GetBookById(id string) models.Book {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookById", id)
	ret0, _ := ret[0].(models.Book)
	return ret0
}

// GetBookById indicates an expected call of GetBookById.
func (mr *MockDbMockRecorder) GetBookById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookById", reflect.TypeOf((*MockDb)(nil).GetBookById), id)
}

// GetBookByTitle mocks base method.
func (m *MockDb) GetBookByTitle(title string) (models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookByTitle", title)
	ret0, _ := ret[0].(models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByTitle indicates an expected call of GetBookByTitle.
func (mr *MockDbMockRecorder) GetBookByTitle(title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByTitle", reflect.TypeOf((*MockDb)(nil).GetBookByTitle), title)
}

// GetBorrowedBooks mocks base method.
func (m *MockDb) GetBorrowedBooks(studentId string) ([]map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBorrowedBooks", studentId)
	ret0, _ := ret[0].([]map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBorrowedBooks indicates an expected call of GetBorrowedBooks.
func (mr *MockDbMockRecorder) GetBorrowedBooks(studentId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBorrowedBooks", reflect.TypeOf((*MockDb)(nil).GetBorrowedBooks), studentId)
}

// GetStudentByName mocks base method.
func (m *MockDb) GetStudentByName(first, last string) (models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudentByName", first, last)
	ret0, _ := ret[0].(models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudentByName indicates an expected call of GetStudentByName.
func (mr *MockDbMockRecorder) GetStudentByName(first, last interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentByName", reflect.TypeOf((*MockDb)(nil).GetStudentByName), first, last)
}

// ReturnBook mocks base method.
func (m *MockDb) ReturnBook(studentId, bookId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturnBook", studentId, bookId)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReturnBook indicates an expected call of ReturnBook.
func (mr *MockDbMockRecorder) ReturnBook(studentId, bookId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReturnBook", reflect.TypeOf((*MockDb)(nil).ReturnBook), studentId, bookId)
}

// StudentCheckIn mocks base method.
func (m *MockDb) StudentCheckIn(s models.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StudentCheckIn", s)
	ret0, _ := ret[0].(error)
	return ret0
}

// StudentCheckIn indicates an expected call of StudentCheckIn.
func (mr *MockDbMockRecorder) StudentCheckIn(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StudentCheckIn", reflect.TypeOf((*MockDb)(nil).StudentCheckIn), s)
}

// StudentLogin mocks base method.
func (m *MockDb) StudentLogin(username, password string) (models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StudentLogin", username, password)
	ret0, _ := ret[0].(models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StudentLogin indicates an expected call of StudentLogin.
func (mr *MockDbMockRecorder) StudentLogin(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StudentLogin", reflect.TypeOf((*MockDb)(nil).StudentLogin), username, password)
}

// StudentSignUp mocks base method.
func (m *MockDb) StudentSignUp(s models.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StudentSignUp", s)
	ret0, _ := ret[0].(error)
	return ret0
}

// StudentSignUp indicates an expected call of StudentSignUp.
func (mr *MockDbMockRecorder) StudentSignUp(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StudentSignUp", reflect.TypeOf((*MockDb)(nil).StudentSignUp), s)
}

// UpdateBook mocks base method.
func (m *MockDb) UpdateBook(book models.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", book)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockDbMockRecorder) UpdateBook(book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockDb)(nil).UpdateBook), book)
}

// UpdateBookStatus mocks base method.
func (m *MockDb) UpdateBookStatus(status bool, bookID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBookStatus", status, bookID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBookStatus indicates an expected call of UpdateBookStatus.
func (mr *MockDbMockRecorder) UpdateBookStatus(status, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBookStatus", reflect.TypeOf((*MockDb)(nil).UpdateBookStatus), status, bookID)
}

// UpdateStockCount mocks base method.
func (m *MockDb) UpdateStockCount(bookID, stock string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStockCount", bookID, stock)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStockCount indicates an expected call of UpdateStockCount.
func (mr *MockDbMockRecorder) UpdateStockCount(bookID, stock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStockCount", reflect.TypeOf((*MockDb)(nil).UpdateStockCount), bookID, stock)
}
