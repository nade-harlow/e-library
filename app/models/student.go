package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (db *DbInstance) StudentCheckIn(s Student) error {
	s.ID = uuid.NewString()
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
	stmt, err := db.Postgres.Prepare(fmt.Sprintf("INSERT INTO student_books(id, student_id, book_id, returned, created_at, updated_at) VALUES ($1, $2, $3,$4, $5, $6)"))
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
	stmt, err := db.Postgres.Prepare(fmt.Sprintf("UPDATE student_books SET returned = $1, updated_at = $2 WHERE student_id = $3 AND book_id = $4"))
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
