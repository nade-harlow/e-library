package models

type Book struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Url        string `json:"url"`
	StockCount int    `json:"stock_count"` //
	Available  bool   `json:"available"`
	CreatedAt  string `json:"created_at"`
	ModifiedAt string `json:"modified_at"`
}

type Student struct {
	ID         string `json:"id"`
	UserName   string `json:"user_name"` //
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Password   string `json:"password"` //
	CreatedAt  string `json:"created_at"`
	ModifiedAt string `json:"modified_at"`
}

type BorrowedBook struct {
	ID         string `json:"id"`
	StudentID  string `json:"student_id"`
	BookID     string `json:"book_id"`
	Returned   bool   `json:"returned"`
	CreatedAt  string `json:"created_at"`
	ModifiedAt string `json:"modified_at"`
}
