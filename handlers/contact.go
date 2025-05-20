package handlers

import (
	"fmt"
	"net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		message := r.FormValue("message")
		fmt.Fprintf(w, "Thanks %s, we got your message: %s", name, message)
		return
	}

	fmt.Fprintln(w, `
		<form method="POST" action="/contact">
			Name: <input name="name"><br>
			Message: <textarea name="message"></textarea><br>
			<input type="submit">
		</form>
	`)
}
