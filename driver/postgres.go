package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var db *sql.DB
var err error

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println("Waiting for database connection...")
	time.Sleep(5 * time.Second)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}
	log.Println("Connected to the database successfully")
}

func GetDB() *sql.DB {
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}
	return db
}

func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Println("Error closing the database connection:", err)
		} else {
			log.Println("Database connection closed")
		}
	}
}
