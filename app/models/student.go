package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"
)

func (db *DbInstance) StudentSignUp(s Student) error {
	s.FirstName = strings.ToLower(s.FirstName)
	s.LastName = strings.ToLower(s.LastName)
	stm, err := db.Postgres.Prepare(fmt.Sprintf("INSERT INTO students (id, first_name, last_name, username, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"))
	if err != nil {
		return err
	}
	_, err = stm.Exec(s.ID, s.FirstName, s.LastName, s.UserName, s.Password, time.Now().String(), time.Now().String())
	if err != nil {
		return err
	}
	return err
}

func (db DbInstance) StudentLogin(username, password string) (Student, error) {
	student := Student{}
	row := db.Postgres.QueryRow(fmt.Sprintf("SELECT id, username, first_name FROM students WHERE username = $1 AND password = $2"), username, password)
	err := row.Scan(&student.ID, &student.UserName, &student.FirstName)
	if err != nil {
		log.Println(err.Error())
		return student, err
	}
	return student, nil
}

func (db *DbInstance) StudentCheckIn(s Student) error {
	s.FirstName = strings.ToLower(s.FirstName)
	s.LastName = strings.ToLower(s.LastName)
	stm, err := db.Postgres.Prepare(fmt.Sprintf("INSERT INTO students (id, first_name, last_name, user_name, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"))
	if err != nil {
		return err
	}
	_, err = stm.Exec(s.ID, s.FirstName, s.LastName, s.UserName, s.Password, time.Now().String(), time.Now().String())
	if err != nil {
		return err
	}
	return err
}

func (db DbInstance) GetStudentByName(first, last string) (Student, error) {
	student := Student{}
	row := db.Postgres.QueryRow(fmt.Sprintf("SELECT id, first_name, last_name, user_name, created_at FROM students WHERE first_name = $1 AND last_name = $2"), first, last)
	err := row.Scan(&student.ID, &student.FirstName, &student.LastName, &student.UserName, &student.CreatedAt)
	if err != nil {
		log.Println(err.Error())
		return student, err
	}
	return student, nil
}

func (db DbInstance) GetStudentByUserName(userName string) (Student, error) {
	student := Student{}
	row := db.Postgres.QueryRow(fmt.Sprintf("SELECT id, first_name, last_name, user_name, created_at FROM students WHERE user_name = $1"), userName)
	err := row.Scan(&student.ID, &student.FirstName, &student.LastName, &student.UserName, &student.CreatedAt)
	if err != nil {
		log.Println(err.Error())
		return student, err
	}
	return student, nil
}

func (db *DbInstance) BorrowBook(bookId, studentId string) error {
	stock := db.CheckStockCount(bookId)
	if stock == 0 {
		return errors.New(fmt.Sprintf(`book is no longer in shelf`))
	}
	bbb, _ := db.CheckIfBorrowed(studentId, bookId)
	if bbb.BookID != "" {
		return errors.New("you've already borrowed this book")
	}
	log.Println(bbb.BookID)
	stmt, err := db.Postgres.Prepare(fmt.Sprintf("INSERT INTO borrowed_books(id, student_id, book_id, returned, created_at, updated_at) VALUES ($1, $2, $3,$4, $5, $6)"))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uuid.NewString(), studentId, bookId, false, time.Now().String(), time.Now().String())
	if err != nil {
		return err
	}
	_ = db.UpdateStockCount(bookId, "stock-1")

	return err
}

func (db DbInstance) CheckIfBorrowed(studentID, bookID string) (BorrowedBook, error) {
	bb := BorrowedBook{}
	row := db.Postgres.QueryRow(fmt.Sprintf("SELECT * FROM borrowed_books WHERE student_id = $1 AND book_id = $2 AND returned = false"), studentID, bookID)
	err := row.Scan(&bb.ID, &bb.StudentID, &bb.BookID, &bb.Returned, &bb.CreatedAt, &bb.ModifiedAt)
	if err != nil {
		log.Println(err.Error())
		return bb, err
	}
	return bb, nil
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
	row, err := db.Postgres.Query(fmt.Sprintf("SELECT s.id, s.first_name, s.last_name, k.id, k.title, k.author, k.url FROM (borrowed_books b INNER JOIN students s ON b.student_id = s.id) INNER JOIN books k ON k.id = b.book_id AND b.student_id = $1 AND b.returned = false"), studentId)
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
