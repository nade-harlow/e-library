package models

import (
	"errors"
	"fmt"
)

// GetAllLendings TODO: change student_books to borrowed_books
func (db *DbInstance) GetAllLending() ([]StudentBook, error) {
	var books []StudentBook
	row, err := db.Postgres.Query(fmt.Sprintf("SELECT * FROM student_books"))
	if err != nil {
		return nil, err
	}
	for row.Next() {
		book := StudentBook{}
		err = row.Scan(&book.ID, &book.StudentID, &book.BookID, &book.Returned, &book.CreatedAt, &book.ModifiedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (db *DbInstance) CheckLendStatus(studentId, bookId string) error {
	sb := StudentBook{}
	row := db.Postgres.QueryRow(fmt.Sprintf("SELECT returned FROM student_books WHERE student_id = $1 AND book_id = $2"), studentId, bookId)
	err := row.Scan(&sb.Returned)
	if err != nil {
		return err
	}
	if sb.Returned != true {
		return errors.New("book has not been returned")
	}
	return nil
}
