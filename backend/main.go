package main

import (
	"fmt"
	"net/http"
		"database/sql"
	"log"
	"os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(){
	dsn := os.Getenv("DB_URL")
	var err error 
	
	DB,err = sql.Open("postgres",dsn)
	if err != nil {
		log.Fatalf("Error opening DB: %v",err)
	}

	// Optional: ping to verify connection 
	err = DB.Ping()
	if err != nil {
		log.Fatalf("cannot connect to DB: %v",err)
	}

	fmt.Println("Connected to DB successfully")
}

func main() {
	//connect to db
	InitDB()

	//start server
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	fmt.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", mux)
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a basic hello server")
}
