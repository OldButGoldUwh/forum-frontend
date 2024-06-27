package repository

import (
	"encoding/json"
	"golang-forum-frontend/internal/manager"
	"golang-forum-frontend/internal/models"
	"io"
	"net/http"
)

func GetComments() ([]models.Comment, error) {
	var comments []models.Comment

	apiManager := manager.NewAPIManager()
	apiUrlManager := manager.NewAPIUrls()
	commentsApiUrl := apiUrlManager.GetCommentsApiURL()

	commentsResponse, errComment := apiManager.Get(commentsApiUrl)
	if errComment != nil {
		return nil, errComment
	}

	if commentsResponse.StatusCode == http.StatusOK {
		defer commentsResponse.Body.Close()
		body, err := io.ReadAll(commentsResponse.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &comments)
		if err != nil {
			return nil, err
		}

	}

	return comments, nil
}

func GetPostComments(commentId string) ([]models.Comment, error) {
	var comments []models.Comment
	var commentResponse *http.Response

	apiManager := manager.NewAPIManager()
	apiUrlManager := manager.NewAPIUrls()
	commentApiUrl := apiUrlManager.GetPostsApiURL() + "/" + commentId + "/comments"
	commentResponse, errComment := apiManager.Get(commentApiUrl)
	if commentResponse.StatusCode == http.StatusOK {
		defer commentResponse.Body.Close()
		body, err := io.ReadAll(commentResponse.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &comments)
		if err != nil {
			return nil, err
		}

	}

	if errComment != nil {
		return nil, errComment
	}

	return nil, nil
}
