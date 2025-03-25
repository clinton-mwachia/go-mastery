package main

import (
	"embed"
	"net/http"
)

//go:embed static/*
var content embed.FS

func main() {
	fs := http.FileServer(http.FS(content))
	http.Handle("/static/", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := content.ReadFile("static/index.html")
		if err != nil {
			http.Error(w, "Could not read embedded file", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})

	http.ListenAndServe(":8080", nil)
}
