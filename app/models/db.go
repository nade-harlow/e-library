package models

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

type Db interface {
	Create(book Book) error
	AllBooks() ([]Book, error)
	CheckBookAvailability(title string) (bool, string)
	GetBookById(id string)
	GetBookByTitle(title string) (Book, error)
	ReturnBook(studentId, bookId string) error
	BorrowBook(bookId, studentId string) error
	StudentCheckIn(s Student) error
	CheckLendStatus(studentId, bookId string) error
	GetAllLending() ([]StudentBook, error)
	UpdateBookStatus(status bool, bookID string) error
}

func (db *DbInstance) Create(book Book) error {
	stmt, err := db.Postgres.Prepare(fmt.Sprintf("INSERT INTO books (id, title, author, available, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"))
	if err != nil {
		return err
	}
	book.Title = strings.ToLower(book.Title)
	book.Author = strings.ToLower(book.Author)
	_, err = stmt.Exec(book.ID, book.Title, book.Author, book.Available, book.CreatedAt, book.ModifiedAt)
	if err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) AllBooks() ([]Book, error) {
	books := []Book{}
	row, err := db.Postgres.Query(fmt.Sprintf("SELECT * FROM books"))
	if err != nil {
		return nil, err
	}
	for row.Next() {
		book := Book{}
		err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Available, &book.CreatedAt, &book.ModifiedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (db DbInstance) CheckBookAvailability(title string) (bool, string) {
	book := Book{}
	row := db.Postgres.QueryRow(fmt.Sprintf("SELECT available, author FROM books WHERE title = $1"), title)
	err := row.Scan(&book.Available, &book.Author)
	if err != nil {
		log.Println(err.Error())
		return false, book.Author
	}

	return book.Available, book.Author
}

func (db DbInstance) GetBookById(id string) {
	book := Book{}
	row := db.Postgres.QueryRow(fmt.Sprintf("SELECT * FROM books WHERE id = $1"), id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Available, &book.CreatedAt, &book.ModifiedAt)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (db DbInstance) GetBookByTitle(title string) (Book, error) {
	book := Book{}
	_, exist := db.CheckBookAvailability(title)
	if exist == "" {
		return book, errors.New("book does not exist")
	}
	row := db.Postgres.QueryRow(fmt.Sprintf("SELECT * FROM books WHERE title = $1"), title)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Available, &book.CreatedAt, &book.ModifiedAt)
	if err != nil {
		log.Println(err.Error())
		return book, err
	}
	return book, nil
}

func (db DbInstance) UpdateBookStatus(status bool, bookID string) error {
	stmt, err := db.Postgres.Prepare(fmt.Sprintf("UPDATE books SET available = $1, updated_at = $2 WHERE id = $3"))
	if err != nil {
		return err
	}
	_, er := stmt.Exec(status, time.Now().String(), bookID)
	if er != nil {
		return er
	}
	return nil
}
