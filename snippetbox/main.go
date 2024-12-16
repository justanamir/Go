package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Display the home page.
func home (w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from Snippetbox"))
}

// Display a specific snippet.
func snippetView (w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Display a form for creating a new snippet.
func snippetCreate (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet"))
}

// Save a new snippet.
func snippetCreatePost (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	// Display the home page.
	mux.HandleFunc("GET /{$}", home)
	// Display a specific snippet.
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	// Display a form for creating a new snippet.
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	// Save a new snippet.
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}