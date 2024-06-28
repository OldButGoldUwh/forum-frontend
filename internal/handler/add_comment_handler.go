package handler

import (
	"fmt"
	"golang-forum-frontend/internal/models"
	"golang-forum-frontend/internal/repository"
	"net/http"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Comment Handler")
	if r.Method == http.MethodPost {

		postID := r.FormValue("postID")
		content := r.FormValue("content")
		fmt.Println("Post ID:", postID)  // Debug statement
		fmt.Println("Content:", content) // Debug statement

		if postID == "" {
			http.Error(w, "Post ID eksik", http.StatusBadRequest)
			return
		}

		comment := models.Comment{
			Content: content,
		}

		err := repository.AddComment(postID, comment)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/post/%s", postID), http.StatusSeeOther)
	}
}
