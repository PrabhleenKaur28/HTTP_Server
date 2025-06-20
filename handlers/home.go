package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r) //sends a standard "404 Not Found" response
		return
	}
	fmt.Fprintln(w, "Welcome to GoServe!") //Fprintln is for giving file as output
}
