package models

type Db interface {
	AddBook(book Book) error
	GetAllBooks() ([]Book, error)
	CheckBookAvailability(title string) (bool, string)
	GetBookById(id string)
	GetBookByTitle(title string) (Book, error)
	GetBook(title string) (Book, error)
	ReturnBook(studentId, bookId string) error
	BorrowBook(bookId, studentId string) error
	StudentCheckIn(s Student) error
	CheckLendStatus(studentId, bookId string) error
	GetAllLending() ([]BorrowedBook, error)
	UpdateBookStatus(status bool, bookID string) error
	DeleteBookById(bookID string) error
	DeleteBookByTitle(bookTitle string) error
}
