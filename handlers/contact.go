package handlers

import (
	"fmt"
	"net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "./static/contact.html")
		return
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		// Use all variables
		fmt.Fprintf(w, "Thanks %s!\n\nWe received your message:\n\"%s\"\n\nWe'll contact you at: %s", name, message, email)
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
