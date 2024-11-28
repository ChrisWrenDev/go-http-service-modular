package main

import (
	"net/http"
)

func commentsHandler(store *CommentStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Handle the request
    }
}

func anotherHandler(store *AnotherStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Handle the request
    }
}
