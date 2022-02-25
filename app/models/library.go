package models

type Library struct {
	ID string `json:"id"`
	Books []Book `json:"books"`
	CreatedAt string `json:"created_at"`
	ModifiedAt string `json:"modified_at"`
}

type Book struct {
	ID string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Available bool `json:"available"`
	CreatedAt string `json:"created_at"`
	ModifiedAt string `json:"modified_at"`
}

type Student struct {
	ID string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	BorrowedBook []Book `json:"books"`
	CreatedAt string `json:"created_at"`
	ModifiedAt string `json:"modified_at"`
}

