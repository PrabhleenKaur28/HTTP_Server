package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/PrabhleenKaur28/HTTP_Server/db"
)

// ContactMessage represents one contact entry
type ContactMessage struct {
	ID          int
	Name        string
	Email       string
	Message     string
	SubmittedAt time.Time
}

// MessagesHandler fetches and shows all contact messages
func MessagesHandler(w http.ResponseWriter, r *http.Request) {
	// Query all messages ordered by submission date desc
	rows, err := db.DB.Query(`SELECT id, name, email, message, submitted_at FROM contacts ORDER BY submitted_at DESC`)
	if err != nil {
        showError(w, "Sorry, we're having trouble loading messages. Please try again later.")
        fmt.Println("DB query error:", err)
        return
    }
	defer rows.Close()

	messages := []ContactMessage{}

	for rows.Next() {
		var msg ContactMessage
		err := rows.Scan(&msg.ID, &msg.Name, &msg.Email, &msg.Message, &msg.SubmittedAt)
		if err != nil {
			http.Error(w, "Error reading messages", http.StatusInternalServerError)
			fmt.Println("Row scan error:", err)
			return
		}
		messages = append(messages, msg)
	}
	t, err := template.ParseFiles("templates/messages.html")
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        fmt.Println("Template error:", err)
        return
    }
	err = t.Execute(w, messages)
	if err != nil {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
		fmt.Println("Template exec error:", err)
	}
}

func showError(w http.ResponseWriter, message string) {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Error</title>
            <link rel="stylesheet" href="/static/css/admin.css">
        </head>
        <body>
            <div class="error">%s</div>
        </body>
        </html>
    `, message)
}

func DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    id := r.FormValue("id")
    _, err := db.DB.Exec("DELETE FROM contacts WHERE id = $1", id)
    if err != nil {
        http.Error(w, "Error deleting message", http.StatusInternalServerError)
        return
    }

    // Redirect back to messages page
    http.Redirect(w, r, "/admin/messages", http.StatusSeeOther)
}