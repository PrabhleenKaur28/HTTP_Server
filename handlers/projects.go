package handlers

import (
	"encoding/json"
	"net/http"
)

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects := []Project{
		{Name: "PACER", Description: "P2P CDN"},
		{Name: "Library System", Description: "Book tracking with trie search"},
		{Name: "Snake Game", Description: "JavaScript game with difficulty modes"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
