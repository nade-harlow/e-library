package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"
)

func (db *DbInstance) AddBook(book Book) error {
	book.Title = strings.ToLower(book.Title)
	book.Author = strings.ToLower(book.Author)
	book.ID = uuid.NewString()
	book.Available = true
	book.CreatedAt = time.Now().String()
	book.ModifiedAt = time.Now().String()
	_, exist := db.CheckBookAvailability(book.Title)
	if exist != "" {
		return errors.New(fmt.Sprintf(`'%s' Already exist in library`, book.Title))
	}
	stmt, err := db.Postgres.Prepare(fmt.Sprintf("INSERT INTO books (id, title, author, url, available, stock, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(book.ID, book.Title, book.Author, book.Url, book.Available, book.StockCount, book.CreatedAt, book.ModifiedAt)
	if err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) GetAllBooks() ([]Book, error) {
	books := []Book{}
	row, err := db.Postgres.Query(fmt.Sprintf("SELECT * FROM books ORDER BY created_at ASC"))
	if err != nil {
		return nil, err
	}
	for row.Next() {
		book := Book{}
		err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Url, &book.StockCount, &book.Available, &book.CreatedAt, &book.ModifiedAt)
		if err != nil {
			return nil, err
		}
		book.CreatedAt = book.CreatedAt[:19]
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

func (db DbInstance) GetBookById(bookID string) Book {
	book := Book{}
	row := db.Postgres.QueryRow(fmt.Sprintf("SELECT * FROM books WHERE id = $1"), bookID)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Url, &book.StockCount, &book.Available, &book.CreatedAt, &book.ModifiedAt)
	if err != nil {
		log.Println(err.Error())
		return book
	}
	return book
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
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Url, &book.StockCount, &book.Available, &book.CreatedAt, &book.ModifiedAt)
	if err != nil {
		log.Println(err.Error())
		return book, err
	}
	return book, nil
}

func (db DbInstance) UpdateStockCount(bookID, stock string) error {
	stmt, err := db.Postgres.Prepare(fmt.Sprintf("UPDATE books SET stock = %s, updated_at = $1 WHERE id = $2", stock))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(time.Now().String(), bookID)
	if err != nil {
		return err
	}

	return nil
}

func (db DbInstance) CheckStockCount(bookID string) int {
	book := Book{}
	row := db.Postgres.QueryRow(fmt.Sprintf("SELECT stock FROM books WHERE id = $1"), bookID)
	err := row.Scan(&book.StockCount)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return book.StockCount
}

func (db DbInstance) GetBook(title string) (Book, error) {
	book := Book{}
	row := db.Postgres.QueryRow(fmt.Sprintf("SELECT * FROM books WHERE title = $1"), title)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Url, &book.StockCount, &book.Available, &book.CreatedAt, &book.ModifiedAt)
	if err != nil {
		log.Println(err.Error())
		return book, err
	}
	return book, nil
}

func (db DbInstance) UpdateBook(book Book) error {
	stmt, err := db.Postgres.Prepare(fmt.Sprintf("UPDATE books SET title = $1, author = $2, url = $3, available = $4, stock = $5, updated_at = $6 WHERE id = $7"))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(book.Title, book.Author, book.Url, true, book.StockCount, time.Now().String(), book.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) UpdateBookStatus(status bool, bookID string) error {
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

func (db *DbInstance) DeleteBookById(bookID string) error {
	stmt, err := db.Postgres.Prepare(fmt.Sprintf("DELETE FROM books WHERE id = $1"))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(bookID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) DeleteBookByTitle(bookTitle string) error {
	stmt, err := db.Postgres.Prepare(fmt.Sprintf("DELETE FROM books WHERE title = $1"))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(bookTitle)
	if err != nil {
		return err
	}
	return nil
}
