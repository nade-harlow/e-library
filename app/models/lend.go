package models

import (
	"errors"
	"fmt"
)

func (db *DbInstance) GetAllLending(returned bool) ([]map[string]interface{}, error) {
	var Bbooks []map[string]interface{}
	row, err := db.Postgres.Query(fmt.Sprintf("SELECT s.id, s.first_name, s.last_name, k.id, k.title, k.author, k.url FROM (borrowed_books b INNER JOIN students s ON b.student_id = s.id) INNER JOIN books k ON k.id = b.book_id AND b.returned = $1"), returned)
	if err != nil {
		return nil, err
	}
	for row.Next() {

		var StudentID, FirstName, LastName, BookID, BookTitle, BookAuthor, BookUrl string
		err = row.Scan(&StudentID, &FirstName, &LastName, &BookID, &BookTitle, &BookAuthor, &BookUrl)
		if err != nil {
			return nil, err
		}
		borrowedBooks := map[string]interface{}{
			"StudentID":  StudentID,
			"FirstName":  FirstName,
			"LastName":   LastName,
			"BookID":     BookID,
			"BookTitle":  BookTitle,
			"BookAuthor": BookAuthor,
			"BookUrl":    BookUrl,
		}
		Bbooks = append(Bbooks, borrowedBooks)
	}
	return Bbooks, nil
}

func (db *DbInstance) CheckLendStatus(studentId, bookId string) error {
	sb := BorrowedBook{}
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
