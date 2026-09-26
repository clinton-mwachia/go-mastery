package main

import (
	"go/templates/router"
	"go/templates/utils"
	"log"
	"net/http"
)

type PageData struct {
	Todos []utils.Todo
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	router.RegisterRoutes(mux)

	log.Println("App running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server crash: %v", err)
	}
}
