package middleware

import (
	"log" //logging
	"net/http" //HTTP handling
	"time" //Time functions
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, time.Since(start))
		next(w, r)
	}
}
