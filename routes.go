package main

import (
	"log"
	"net/http"
)

func addRoutes(
    mux *http.ServeMux,
    logger *log.Logger,
    config *Config,
    commentStore *CommentStore,
    anotherStore *AnotherStore,
) {
    mux.HandleFunc("/comments", commentsHandler(commentStore))
    mux.HandleFunc("/another", anotherHandler(anotherStore))
    // Add more routes here
}
