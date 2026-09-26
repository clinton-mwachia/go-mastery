package router

import (
	"go/templates/utils"
	"html/template"
	"net/http"
	"strconv"
)

type PageData struct {
	Todos []utils.Todo
}

func RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		data := PageData{Todos: utils.GetAll()}
		tmpl, err := template.ParseFiles("ui/index.tmpl")
		if err != nil {
			http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	})

	mux.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		if title == "" {
			http.Error(w, "Title cannot be empty", http.StatusBadRequest)
			return
		}
		utils.Add(title)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	mux.HandleFunc("POST /toggle", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}
		if err := utils.Toggle(id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	mux.HandleFunc("POST /delete", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}
		if err := utils.Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}
