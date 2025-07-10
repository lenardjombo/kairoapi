package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lenardjombo/gallus/pkg"
)

func main() {
	// 1. Initialize the DB connection
	if err := pkg.Init(); err != nil {
		log.Fatalf("Failed to connect to the DB: %s", err)
	}

	// 2. Inform the user that the server is starting
	fmt.Println("Server is running on http://127.0.0.1:8080")

	// 3. Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
