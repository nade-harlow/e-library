package models

//go:generate mockgen -destination=../mocks/mock_db.go -package=mocks github.com/nade-harlow/e-library/models Db
type Db interface {
	AddBook(book Book) error
	GetAllBooks() ([]Book, error)
	CheckBookAvailability(title string) (bool, string)
	GetBookById(id string)
	GetBookByTitle(title string) (Book, error)
	GetBook(title string) (Book, error)
	ReturnBook(studentId, bookId string) error
	BorrowBook(bookId, studentId string) error
	GetBorrowedBooks(studentId string) ([]map[string]interface{}, error)
	StudentCheckIn(s Student) error
	StudentSignUp(s Student) error
	CheckIfBorrowed(studentID, bookID string) (BorrowedBook, error)
	GetStudentByName(first, last string) (Student, error)
	CheckLendStatus(studentId, bookId string) error
	GetAllLending(returned bool) ([]map[string]interface{}, error)
	UpdateBookStatus(status bool, bookID string) error
	DeleteBookById(bookID string) error
	DeleteBookByTitle(bookTitle string) error
}
