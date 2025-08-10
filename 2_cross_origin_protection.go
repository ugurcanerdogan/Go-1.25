package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Registering handlers
	mux := http.NewServeMux()
	mux.HandleFunc("GET /get", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "GET ok\n")
	})
	mux.HandleFunc("POST /post", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "POST ok\n")
	})

	// Also:
	// handler := http.CrossOriginProtection(mux)
	// http.ListenAndServe(":8080", handler)

	// Configuration for CSRF protection
	antiCSRF := http.NewCrossOriginProtection()
	antiCSRF.AddTrustedOrigin("https://trendyol.com")
	antiCSRF.AddTrustedOrigin("https://*.trendyol.com")

	// Adding CSRF protection to all handlers
	srv := http.Server{
		Addr:    ":8080",
		Handler: antiCSRF.Handler(mux),
	}
	log.Fatal(srv.ListenAndServe())
}
