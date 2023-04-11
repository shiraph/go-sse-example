package main

import (
	"log"
	"net/http"
	"time"

	"github.com/guregu/null"
	"github.com/r3labs/sse/v2"
	"github.com/rs/cors"

	"sse-example/internal/event_server"
	"sse-example/internal/watcher"
)

const (
	DefaultPath            = "/"
	Addr                   = ":3000"
	StreamName             = "messages"
	DefaultMessageInterval = 1 * time.Second
	MessageLimit           = false
	NumberOfMessages       = 10 // only works when MessageLimit is true
)

func main() {
	// Create a new event server.
	eventServer := event_server.NewEventServer(sse.New(), StreamName, DefaultMessageInterval)
	eventServer.Init(null.NewInt(NumberOfMessages, MessageLimit).Ptr())

	// Create a new HTTP mux and set the handler.
	mux := http.NewServeMux()
	mux.HandleFunc(DefaultPath, func(w http.ResponseWriter, r *http.Request) {
		go func() {
			<-r.Context().Done()
		}()
		eventServer.Server.ServeHTTP(w, r)
	})

	// Create a new watcher and server.
	var w watcher.Watcher
	c := cors.AllowAll()
	server := &http.Server{
		Addr:      Addr,
		Handler:   c.Handler(mux),
		ConnState: w.OnStateChange,
	}

	// Start the server.
	log.Printf("Starting server on %s", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
