package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
    logger := log.New(os.Stdout, "", log.LstdFlags)
    config := &Config{
        Host: "localhost",
        Port: "8080",
    }
    // Initialize your stores here
    commentStore := &CommentStore{}
    anotherStore := &AnotherStore{}

    srv := NewServer(logger, config, commentStore, anotherStore)
    httpServer := &http.Server{
        Addr:    net.JoinHostPort(config.Host, config.Port),
        Handler: srv,
    }

    go func() {
        logger.Printf("Listening on %s\n", httpServer.Addr)
        if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Fatalf("Error listening and serving: %s\n", err)
        }
    }()

    // Graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    logger.Println("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := httpServer.Shutdown(ctx); err != nil {
        logger.Fatalf("Server forced to shutdown: %s\n", err)
    }

    logger.Println("Server exiting")
}
