package main

import (
	"log"
	"net/http"
)

var artistsCache []ArtistView

func main() {
	err := LoadArtists()
	if err != nil {
		log.Fatal("Failed to load data:", err)
	}
	http.HandleFunc("/static/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/style.css")
	})
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artist", artistHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
