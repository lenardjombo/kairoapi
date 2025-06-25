package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL environment variable is required")
	}

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	// Verify connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	fmt.Println("Connected to DB successfully")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}