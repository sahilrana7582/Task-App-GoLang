package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() *sql.DB {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	fmt.Println("DB Connection String:", connStr) // TEMPORARY

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Failed to open DB: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	log.Println("Connected to DB successfully")

	DB = db

	return db

}
