package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the About page.")
}

func main() {
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/about", aboutHandler)

    fmt.Println("Server is starting at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
