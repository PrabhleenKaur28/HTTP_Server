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
		http.Error(w, "Error fetching messages", http.StatusInternalServerError)
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

	const tpl = `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Contact Messages</title>
			<style>
				table { border-collapse: collapse; width: 100%; }
				th, td { border: 1px solid #ccc; padding: 8px; }
				th { background: #f4f4f4; }
			</style>
		</head>
		<body>
			<h1>Contact Messages</h1>
			<table>
				<tr>
					<th>ID</th><th>Name</th><th>Email</th><th>Message</th><th>Submitted At</th>
				</tr>
				{{range .}}
					<tr>
						<td>{{.ID}}</td>
						<td>{{.Name}}</td>
						<td>{{.Email}}</td>
						<td>{{.Message}}</td>
						<td>{{.SubmittedAt.Format "2006-01-02 15:04:05"}}</td>
					</tr>
				{{else}}
					<tr><td colspan="5">No messages found</td></tr>
				{{end}}
			</table>
		</body>
		</html>`

	t := template.Must(template.New("messages").Parse(tpl))
	err = t.Execute(w, messages)
	if err != nil {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
		fmt.Println("Template exec error:", err)
	}
}
