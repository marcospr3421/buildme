package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// version is defined using build args
var version string

func main() {
	if version != "" {
		log.Printf("Version: %s", version)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
	// Serve homepage with embedded image
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Docker Sample Image</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					text-align: center;
					margin-top: 50px;
				}
				.container {
					max-width: 800px;
					margin: 0 auto;
				}
				img {
					max-width: 100%;
					height: auto;
					border: 1px solid #ddd;
					border-radius: 4px;
					padding: 5px;
				}
				h1 {
					color: #333;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Welcome to Your Docker Container!</h1>
				<p>This is a sample image being served from your Docker container</p>
				<img src="data:image/png;base64,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" alt="Docker Logo">
			</div>
		</body>
		</html>
		`
		fmt.Fprint(w, html)
	})
	
	// Keep the original API endpoint for translation
	r.Get("/translate", func(w http.ResponseWriter, r *http.Request) {
		msg := r.URL.Query().Get("s")
		translated := translate(msg)
		w.Write([]byte(translated))
	})

	log.Println("Starting server...")
	log.Println("Listening on HTTP port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
