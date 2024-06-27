package handler

import (
	"encoding/json"
	"fmt"
	"golang-forum-frontend/internal/models"
	"golang-forum-frontend/internal/repository"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	layout := filepath.Join("web", "templates", "layout.html")
	navbar := filepath.Join("web", "templates", "navbar.html")
	search := filepath.Join("web", "templates", "search.html")
	postHtml := filepath.Join("web", "templates", "post.html")

	vars := mux.Vars(r)
	postID := vars["id"]
	var post models.Post
	var comments []models.Comment

	post, errPost := repository.GetPost(postID)
	comments, errComment := repository.GetPostComments(postID)

	if errComment != nil {
		http.Error(w, errComment.Error(), http.StatusInternalServerError)
		return
	}

	if errPost != nil {
		http.Error(w, errPost.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(layout, navbar, search, postHtml)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username, err := repository.GetUsername()

	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Username:", username) // Log the username

	data := struct {
		Post     models.Post
		Comments []models.Comment
		Username string
	}{
		Post:     post,
		Comments: comments,
		Username: username,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetPostCommentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]
	comments, err := repository.GetPostComments(postID)
	fmt.Println("Comments:", comments) // Log the comments
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
