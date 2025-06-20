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
	defer rows.Close()//schedules the rows to be closed when func exits to prevent resource leaks
	//defer ensures cleanup happens no matter how the func exits: normally, early return or panic. Hence, better than manually closing at each return point
	//Databases have a limited number of simultaneous connections they can handle. When we execute a query, it uses one of these connections. Closing the rows releases the connection back to the pool. If we forget to close rows, connections remain occupied, and eventually the database will refuse new requests when all connections are exhausted.

	messages := []ContactMessage{}//creates an empty slice[] (arraylist(flexible)) to hold messages

	for rows.Next() {//returns false if no more rows
		var msg ContactMessage
		err := rows.Scan(&msg.ID, &msg.Name, &msg.Email, &msg.Message, &msg.SubmittedAt)// rows.Scan() reads columns from current row and &msg.Field provides memory addresses to store each value
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