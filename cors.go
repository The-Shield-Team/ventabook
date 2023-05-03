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
		c.Handler(handleFunc).ServeHTTP(w, r)
	}
}
