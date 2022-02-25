package models

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type DbInstance struct {
	Postgres *sql.DB
}

func Init() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := postgresql()
	fmt.Println(db)
	//NewInstance(db)
	return db
}

func postgresql() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("failed to open database %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to verify database connection %s", err.Error())
	}

	log.Println("Successfully connected!")
	return db
}

func NewInstance(db *sql.DB) *DbInstance {
	return &DbInstance{Postgres: db}
}
