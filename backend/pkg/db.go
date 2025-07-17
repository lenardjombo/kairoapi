package pkg

import (
	"database/sql"
	"os"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var  DB *sql.DB

func Init() error{
	// Load .env file 
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("Failed to load .env : %s",err)
	}

	// Connect to DB
	dns := os.Getenv("DB_URI")
	db,err := sql.Open("postgres",dns)
	if err != nil {
		return fmt.Errorf("Error : %s",err)
	}

	DB = db
	//Test Database connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("Error : %s",err)
	}
	return nil
}