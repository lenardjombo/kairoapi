package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	//Load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env-Not Found : %v",err)
	}

	//Connection string
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL environment variable is required")
	}

	//Connect to db
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

// func CloseDB() {
// 	if DB != nil {
// 		DB.Close()
// 	}
// }