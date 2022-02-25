package models

import "log"

type Db interface {
	Create(book Book) error
	AllBooks() ([]Book, error)
}

func (db *DbInstance) Create(book Book) error {
	stmt, err := db.Postgres.Prepare("INSERT INTO book (id, title, author, available, created_at, modified_at) VALUES ($1, $2, $3, $4, $5, $6)")
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
	books:= []Book{}
	row, err := db.Postgres.Query("SELECT * FROM book")
	if err != nil {
		return nil, err
	}
	for row.Next(){
		book:= Book{}
		err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Available, &book.CreatedAt, &book.ModifiedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (db DbInstance) CheckBookAvailability(id string) bool {
	book:= Book{}
	row := db.Postgres.QueryRow("SELECT available FROM book WHERE id = ", id)
	err := row.Scan(&book)
	if err != nil {
		log.Println("")
		return false
	}
	return book.Available
}