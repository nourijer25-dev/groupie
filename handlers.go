package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func testApi(url string) bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url) 
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func HandleError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)

	tmpl := template.Must(template.ParseFiles("templates/error.html"))

	data := ErrorData{
		Code:    code,
		Message: message,
	}

	tmpl.Execute(w, data)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

		if !testApi(artistsCache[0].Image) {
			HandleError(w,http.StatusInternalServerError, "internal server error")
			return
		}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.ExecuteTemplate(w, "index.html", artistsCache)
}

func artistHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(w, http.StatusBadRequest, "Invalid artist ID")
		return
	}
	if id < 1 || id > len(artistsCache) {
		HandleError(w, http.StatusNotFound, "Artist not found")
		return
	}
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
		if !testApi(artistsCache[id-1].Image) {
		fmt.Println("url:", artistsCache[id-1].Image)
		HandleError(w, http.StatusInternalServerError, "Failed to fetch image")
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/artist.html"))
	tmpl.ExecuteTemplate(w, "artist.html", artistsCache[id-1])
}
