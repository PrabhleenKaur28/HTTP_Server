package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PrabhleenKaur28/HTTP_Server/db"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "./static/contact.html")
		return

	case http.MethodPost:
		name := strings.TrimSpace(r.FormValue("name"))
		email := strings.TrimSpace(r.FormValue("email"))
		message := strings.TrimSpace(r.FormValue("message"))

		// Simple validation
		if name == "" || email == "" || message == "" {
			http.Error(w, "All fields are required!", http.StatusBadRequest)
			return
		}

		// Insert into PostgreSQL
		query := `INSERT INTO contacts (name, email, message, submitted_at) VALUES ($1, $2, $3, $4)`
		_, err := db.DB.Exec(query, name, email, message, time.Now())//runs the SQL query with the actual values
		//Exec() returns 2 values: Result(After insertion, these many rows were affected), err
		//_ ignores the result(we know result is also returned but we don't need the result here)
		if err != nil {
			fmt.Println("Error inserting into DB:", err)
			http.Error(w, "Failed to save your message. Please try again later.", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "✅ Thank you! Your message has been received.")
		return

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
