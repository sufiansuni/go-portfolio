package main

import (
	"go-portfolio/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	
	handlers.LoadTemplates()
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.IndexPageHandler)
	router.HandleFunc("/about", handlers.AboutPageHandler)

	log.Fatal(http.ListenAndServe(":" + port, router))
}
