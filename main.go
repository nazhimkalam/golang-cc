package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
) // import all packages

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct / slice is something that can hold multiple values like an array or a collection.
var books []Book

// Function to add mock data
func addMockData() {
	author1 := Author{
		Firstname: "John",
		Lastname:  "Doe",
	}

	author2 := Author{
		Firstname: "Steve",
		Lastname:  "Smith",
	}

	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &author1})
	books = append(books, Book{ID: "2", Isbn: "454545", Title: "Book Two", Author: &author2})
}

// Get All Books Service
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Setting the header to json type response
	json.NewEncoder(w).Encode(books)	// Encoding the books slice into json and sending it to the client
}

// Get Single Book Service
func getBook(w http.ResponseWriter, r *http.Request) {}

// Create a new Book
func createBook(w http.ResponseWriter, r *http.Request) {}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {}

// Main function
func main() {

	// Initialize the Router
	router := mux.NewRouter()

	// Adding Mock data (Hard Coded)
	addMockData()

	// Route Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":5000", router))

}
