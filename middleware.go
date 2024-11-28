package main

import (
	"log"
	"net/http"
)

func loggingMiddleware(logger *log.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            logger.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
            next.ServeHTTP(w, r)
        })
    }
}

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check authentication
        next.ServeHTTP(w, r)
    })
}
