package main

import (
	"log"
	"net/http"

	"./routers"
	"github.com/gorilla/handlers"
)

func main() {
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(headers, origins, methods)(routers.Router)))
}
