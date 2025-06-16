package middleware

import (
	"net/http"
)

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Very simple authentication - in real apps, use proper sessions or JWT
		user, pass, ok := r.BasicAuth()
		
		if !ok || user != "admin" || pass != "admin123" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		
		next(w, r)
	}
}