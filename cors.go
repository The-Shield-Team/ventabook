package main

import (
	"net/http"

	"github.com/rs/cors"
)

func corsMiddleware(handleFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Configurar las opciones CORS
		c := cors.New(cors.Options{
			AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"}, // Permitir solo el método HTTP POST
		})

		// Aplicar las opciones CORS a la solicitud
		handler := c.Handler(handleFunc)
		// Agregar los encabezados CORS a la respuesta HTTP
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")

		// Terminar de procesar la solicitud
		handler.ServeHTTP(w, r)
	}
}
