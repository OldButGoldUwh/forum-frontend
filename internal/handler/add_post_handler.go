package handler

import (
	"fmt"
	"golang-forum-frontend/internal/models"
	"golang-forum-frontend/internal/repository"
	"html/template"
	"net/http"
	"path/filepath"
)

func AddPostHandler(w http.ResponseWriter, r *http.Request) {

	layout := filepath.Join("web", "templates", "layout.html")
	navbar := filepath.Join("web", "templates", "navbar.html")
	search := filepath.Join("web", "templates", "search.html")
	addPostHtml := filepath.Join("web", "templates", "add_post_form.html")
	tmpl, err := template.ParseFiles(layout, navbar, search, addPostHtml)

	if err != nil {
		fmt.Println("Error parsing templates:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("Error executing template: 2", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func AddPostSubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")

		post := models.Post{
			Title:   title,
			Content: content,
		}

		err := repository.AddPost(post)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
