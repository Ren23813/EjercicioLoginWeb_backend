package main

import (
    "net/http"

    "github.com/go-chi/cors"
)


func configureCORS() func(http.Handler) http.Handler {
    corsMiddleware := cors.New(cors.Options{
        // Permitir orígenes específicos (ej. donde corre su frontend)
        AllowedOrigins:   []string{"*", "http://localhost:5500", "http://127.0.0.1:5500"}, 
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true, // Importante si usan cookies o auth headers
        MaxAge:           300, // Maximum value not ignored by any of major browsers
    })
	return corseMiddlware.Handler
}
