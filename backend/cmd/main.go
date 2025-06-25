package main

import (
	"fmt"
	"net/http"

	"github.com/lenardjombo/gallus/pkg"
)

func main() {
	//connect to db
	pkg.InitDB()
	defer pkg.CloseDB()

	//start server
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	fmt.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", mux)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "sample gallus server")
}
