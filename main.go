package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	router.Use(c.Handler)
	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/{params:.+}", getBookByAny).Methods("GET")
	router.HandleFunc("/facturar", Facturar).Methods("POST")

	router.HandleFunc("/books/{id}/{stock}/{bodega}", updateOneBook).Methods("PUT")
	router.HandleFunc("/", sayHello).Methods("GET")

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
