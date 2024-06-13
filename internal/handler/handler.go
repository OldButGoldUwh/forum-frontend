package handler

import (
	"encoding/json"
	"golang-forum-frontend/internal/manager"
	"html/template"
	"net/http"
	"path/filepath"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	layout := filepath.Join("web", "templates", "layout.html")
	index := filepath.Join("web", "templates", "index.html")
	navbar := filepath.Join("web", "templates", "navbar.html")
	search := filepath.Join("web", "templates", "search.html")
	post := filepath.Join("web", "templates", "post.html")

	tmpl, err := template.ParseFiles(layout, index, navbar, search, post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	layout := filepath.Join("web", "templates", "layout.html")
	navbar := filepath.Join("web", "templates", "navbar.html")
	login := filepath.Join("web", "templates", "login.html")
	search := filepath.Join("web", "templates", "search.html")

	tmpl, err := template.ParseFiles(layout, navbar, login, search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = tmpl.Execute(w, nil)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	layout := filepath.Join("web", "templates", "layout.html")
	navbar := filepath.Join("web", "templates", "navbar.html")
	register := filepath.Join("web", "templates", "register.html")
	search := filepath.Join("web", "templates", "search.html")

	email := r.FormValue("email")
	password := r.FormValue("password")
	username := r.FormValue("username")

	user := User{
		Email:    email,
		Password: password,
		Username: username,
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "JSON verisi oluşturulurken hata oluştu", http.StatusInternalServerError)
		return
	}

	//var apiManager manager.APIManager
	apiManager := manager.NewAPIManager()
	apiURL := "http://localhost:8080/api/v1/users"
	apiManager.Post(apiURL, jsonData)

	tmpl, err := template.ParseFiles(layout, navbar, register, search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = tmpl.Execute(w, nil)
}
