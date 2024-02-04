package main

import (
	_ "ApiServer/docs"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

var tokenAuth *jwtauth.JWTAuth
var users = make(map[string]string) // Мапа для хранения пользователей и их паролей

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

// @title Топовый API
// @version 3.0
// @description Предаставляет данные по названию или гео координатам
// @host localhost:8080
func main() {
	addr := ":8080"
	fmt.Printf("Starting server on %v\n", addr)
	err := http.ListenAndServe(addr, router())
	if err != nil {
		log.Fatal(err)
	}
}

func router() http.Handler {
	r := chi.NewRouter()

	// Public routes
	r.Group(func(r chi.Router) {
		r.Post("/api/login", loginHandler)
		r.Post("/api/register", registerHandler)
		r.Get("/swagger/*", httpSwagger.WrapHandler)
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens
		r.Use(jwtauth.Authenticator)

		r.Post("/api/address/search", AddressSearchHandler)
		r.Post("/api/address/geocode", AddressGeocodeHandler)
	})

	return r
}
