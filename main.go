package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Analyser l'URL demandée
		url := r.URL.Path

		// Renvoyer le fichier HTML correspondant en fonction de l'URL demandée
		switch {
		case strings.Contains(url, "portfolio"):
			http.ServeFile(w, r, "portfolio.html")
		case strings.Contains(url, "forum"):
			http.ServeFile(w, r, "forum.html")
		case strings.Contains(url, "projet_infra"):
			http.ServeFile(w, r, "projet_infra.html")
		default:
			http.ServeFile(w, r, "portfolio.html") // Par défaut, servir portfolio.html
		}
	})

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
