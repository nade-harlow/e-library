package models

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
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
	//psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", "postgres://evrhmwwrktikrl:6f6ca098e07d85aa7250a7397d06f79dc0d06e120b99a1c1d7d8ac2432118067@ec2-54-195-76-73.eu-west-1.compute.amazonaws.com:5432/dfemrlacff0t1")
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
