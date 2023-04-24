package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// type con la información de los libros

type book struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Place string `json:"Place"`
}

type allBook []book

var books = allBook{
	{
		ID:    1,
		Name:  "El Quijote xD",
		Place: "Bodega 1"},
	{

		ID:    1,
		Name:  "El Quijote xD",
		Place: "Bodega 2"},
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Holi")
}

// get method
func methodsBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	} else if r.Method == "POST" {
		var newBook book
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Inserte libro")
		}
		json.Unmarshal(reqBody, &newBook)
		newBook.ID = len(books) + 1
		books = append(books, newBook)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newBook)
		w.WriteHeader(http.StatusCreated)

	}

}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/books", methodsBook)
	// http.HandleFunc("/books", createBook).Methods("POST")

	// puerto donde vere el WS
	http.ListenAndServe(":3250", nil)
}
