package main

import (
	"log"
	"net/http"
)

type Server struct {
    mux *http.ServeMux
}

func NewServer(
    logger *log.Logger,
    config *Config,
    commentStore *CommentStore,
    anotherStore *AnotherStore,
) http.Handler {
    mux := http.NewServeMux()
    addRoutes(mux, logger, config, commentStore, anotherStore)

    var handler http.Handler = mux
    handler = loggingMiddleware(logger)(handler)
    handler = authMiddleware(handler)

    return handler
}
