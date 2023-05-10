package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/books", corsMiddleware(getAllBooks)).Methods("GET")
	router.HandleFunc("/books/{params:.+}", corsMiddleware(getBookByAny)).Methods("GET")
	router.HandleFunc("/facturar", corsMiddleware(Facturar)).Methods("POST")

	router.HandleFunc("/books/{id}/{stock}/{bodega}", corsMiddleware(updateOneBook)).Methods("PUT")
	router.HandleFunc("/", corsMiddleware(sayHello)).Methods("GET")
	router.HandleFunc("/last", corsMiddleware(Lastito)).Methods("GET")

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
