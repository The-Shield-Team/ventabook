package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getOneBook).Methods("GET")
	router.HandleFunc("/books/{id}/{stock}/{bodega}", updateOneBook).Methods("PUT")

	// Crear el servidor HTTP
	server := &http.Server{
		Handler:      router,
		Addr:         ":3250",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Iniciar el servidor
	log.Fatal(server.ListenAndServe())

}
