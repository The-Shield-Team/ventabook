package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/{params:.+}", getBookByAny).Methods("GET")
	router.HandleFunc("/facturar", Facturar).Methods("POST")

	router.HandleFunc("/books/{id}/{stock}/{bodega}", updateOneBook).Methods("PUT")
	router.HandleFunc("/", sayHello).Methods("GET")

	// Iniciar el servidor
	log.Fatal(http.ListenAndServe(":3250",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"*"}))(router)))

}
