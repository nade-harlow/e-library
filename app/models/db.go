package models

import (
	"fmt"
	"log"
)

type Db interface {
	Create(book Book) error
	AllBooks() ([]Book, error)
	CheckBookAvailability(id string) bool
	GetBookById(id string)
}

func (db *DbInstance) Create(book Book) error {
	stmt, err := db.Postgres.Prepare(fmt.Sprintf("INSERT INTO book (id, title, author, available, created_at, modified_at) VALUES ($1, $2, $3, $4, $5, $6)"))
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
	row, err := db.Postgres.Query(fmt.Sprintf("SELECT * FROM book"))
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

func (db DbInstance) CheckBookAvailability(id string) bool {
	book := Book{}
	row := db.Postgres.QueryRow(fmt.Sprintf("SELECT available FROM book WHERE id = $1"), id)
	err := row.Scan(&book.Available)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return book.Available
}

func (db DbInstance) GetBookById(id string) {
	book := Book{}
	row := db.Postgres.QueryRow(fmt.Sprintf("SELECT * FROM book WHERE id = $1"), id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Available, &book.CreatedAt, &book.ModifiedAt)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(book)
}

func (db DbInstance) GetBookByTitle(title string) {
	book := Book{}
	row := db.Postgres.QueryRow(fmt.Sprintf("SELECT * FROM book WHERE title = $1"), title)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Available, &book.CreatedAt, &book.ModifiedAt)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(book)
}
