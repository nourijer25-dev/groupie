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
	http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(handleStatic)))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artist", artistHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
