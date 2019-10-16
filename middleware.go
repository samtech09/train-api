package main

import (
	"log"
	"net/http"
)

//loggerMiddleware logs each request to console
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Request for: %s from %s\n", r.URL, r.RemoteAddr)

		next.ServeHTTP(w, r)

	})
}
