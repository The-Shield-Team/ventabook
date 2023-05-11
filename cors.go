package main

import (
	"net/http"

	"github.com/rs/cors"
)

func handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusOK)
}

func corsMiddleware(handleFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Verificar si es una solicitud OPTIONS
		if r.Method == "OPTIONS" {
			handleOptions(w, r)
			return
		}

		// Configurar las opciones CORS
		c := cors.New(cors.Options{
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			Debug:          true,
		})

		// Aplicar las opciones CORS a la solicitud
		handler := c.Handler(handleFunc)

		// Agregar los encabezados CORS a la respuesta HTTP
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Terminar de procesar la solicitud
		handler.ServeHTTP(w, r)
	}
}
