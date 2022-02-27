package models

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

func (db *DbInstance) Create(book Book) error {
	book.Title = strings.ToLower(book.Title)
	book.Author = strings.ToLower(book.Author)
	_, exist := db.CheckBookAvailability(book.Title)
	if exist != "" {
		return errors.New(fmt.Sprintf(`'%s' Already exist in library`, book.Title))
	}
	stmt, err := db.Postgres.Prepare(fmt.Sprintf("INSERT INTO books (id, title, author, available, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"))
	if err != nil {
		return err
	}
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
	available, exist := db.CheckBookAvailability(title)
	if exist == "" {
		return book, errors.New(fmt.Sprintf(`'%s' does not exist in library`, title))
	}
	if !available {
		return book, errors.New(fmt.Sprintf(`'%s' is no longer in shelf`, title))
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

// TODO: delete book method