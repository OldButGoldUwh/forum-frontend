package handler

import (
	"encoding/json"
	"fmt"
	"golang-forum-frontend/internal/manager"
	"golang-forum-frontend/internal/models"
	"golang-forum-frontend/internal/repository"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	layout := filepath.Join("web", "templates", "layout.html")
	index := filepath.Join("web", "templates", "index.html")
	navbar := filepath.Join("web", "templates", "navbar.html")
	search := filepath.Join("web", "templates", "search.html")
	post := filepath.Join("web", "templates", "posts.html")

	tmpl, err := template.ParseFiles(layout, index, navbar, search, post)
	if err != nil {
		fmt.Println("Error parsing templates:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, _ := r.Cookie("token")

	var username string
	var posts []models.Post

	apiManager := manager.NewAPIManager()
	apiUrlManager := manager.NewAPIUrls()

	if token != nil {
		apiManager.SetUserToken(token.Value)

		apiURL := apiUrlManager.GetUserApiURL()

		userResponse, err := apiManager.Get(apiURL)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if userResponse.StatusCode == http.StatusOK {
			defer userResponse.Body.Close()
			body, err := io.ReadAll(userResponse.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var user models.User
			err = json.Unmarshal(body, &user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			username = user.Username

		}

	}

	posts, _ = repository.GetPosts()
	data := struct {
		Username string
		Posts    []models.Post
	}{
		Username: username,
		Posts:    posts,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
