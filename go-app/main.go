package main

import (
	"fmt"
	"net/http"
)

func server(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Go my server!")
}

func test(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Go Test Server!")
}

func main() {
    http.HandleFunc("/", server)
    http.HandleFunc("/test", test)
    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8080", nil)
}
