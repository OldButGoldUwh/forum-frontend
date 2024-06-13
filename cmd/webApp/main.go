package main

import (
	"github.com/gorilla/mux"
	"golang-forum-frontend/internal/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// File server for serving static files
	fileServer := http.FileServer(http.Dir("./web/static/"))
	router.PathPrefix("/web/static/").Handler(http.StripPrefix("/web/static", fileServer))

	// Route for the home handler
	router.HandleFunc("/", handler.HomeHandler)
	router.HandleFunc("/login.html", handler.LoginHandler)
	router.HandleFunc("/register.html", handler.RegisterHandler)
	router.HandleFunc("/register", handler.RegisterHandler)

	// Starting the server
	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
