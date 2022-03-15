package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (db *DbInstance) StudentCheckIn(s Student) error {

	stm, err := db.Postgres.Prepare(fmt.Sprintf("INSERT INTO students (id, first_name, last_name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"))
	if err != nil {
		return err
	}
	_, err = stm.Exec(s.ID, s.FirstName, s.LastName, time.Now().String(), time.Now().String())
	if err != nil {
		return err
	}
	return err
}

func (db *DbInstance) BorrowBook(bookId, studentId string) error {
	stmt, err := db.Postgres.Prepare(fmt.Sprintf("INSERT INTO borrowed_books(id, student_id, book_id, returned, created_at, updated_at) VALUES ($1, $2, $3,$4, $5, $6)"))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uuid.NewString(), studentId, bookId, false, time.Now().String(), time.Now().String())
	if err != nil {
		return err
	}
	return err
}

func (db *DbInstance) ReturnBook(studentId, bookId string) error {
	stmt, err := db.Postgres.Prepare(fmt.Sprintf("UPDATE borrowed_books SET returned = $1, updated_at = $2 WHERE student_id = $3 AND book_id = $4"))
	if err != nil {
		return err
	}
	result, er := stmt.Exec(true, time.Now().String(), studentId, bookId)
	if er != nil {
		return er
	}
	num, _ := result.RowsAffected()
	if num < 1 {
		return errors.New("error updating row")
	}
	return nil
}

func (db DbInstance) GetBorrowedBooks(studentId string) ([]map[string]interface{}, error) {
	var Bbooks []map[string]interface{}
	row, err := db.Postgres.Query(fmt.Sprintf("SELECT s.id, s.first_name, s.last_name, k.id, k.title, k.author FROM (borrowed_books b INNER JOIN students s ON b.student_id = s.id) INNER JOIN books k ON k.id = b.book_id AND b.student_id = $1"), studentId)
	if err != nil {
		return nil, err
	}
	for row.Next() {

		var StudentID, FirstName, LastName, BookID, BookTitle, BookAuthor string
		err = row.Scan(&StudentID, &FirstName, &LastName, &BookID, &BookTitle, &BookAuthor)
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
		}
		Bbooks = append(Bbooks, borrowedBooks)
	}
	return Bbooks, nil
}
