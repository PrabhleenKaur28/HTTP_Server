package main

import (
	"fmt"
	"net/http"

	"github.com/PrabhleenKaur28/HTTP_Server/handlers"
	"github.com/PrabhleenKaur28/HTTP_Server/middleware"
)

func main() {
	mux := http.NewServeMux()

	// Static files
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	// API routes
	mux.HandleFunc("/", middleware.Logger(handlers.HomeHandler))
	mux.HandleFunc("/api/projects", middleware.Logger(handlers.ProjectsHandler))
	mux.HandleFunc("/contact", middleware.Logger(handlers.ContactHandler))

	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
