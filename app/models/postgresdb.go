package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
	return db
}

func NewInstance(db *sql.DB) *DbInstance {
	return &DbInstance{Postgres: db}
}

func postgresql() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	if os.Getenv("DB_USER") == "" {
		db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
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
