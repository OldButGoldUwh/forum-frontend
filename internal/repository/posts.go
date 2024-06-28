package repository

import (
	"encoding/json"
	"fmt"
	"golang-forum-frontend/internal/manager"
	"golang-forum-frontend/internal/models"
	"io"
	"net/http"
)

func GetPosts() ([]models.Post, error) {
	var posts []models.Post

	apiManager := manager.NewAPIManager()
	apiUrlManager := manager.NewAPIUrls()
	postApiUrl := apiUrlManager.GetPostsApiURL()
	fmt.Println("Post API URL:", postApiUrl) // Log the post API URL

	postResponse, errPost := apiManager.Get(postApiUrl)
	if errPost != nil {

		return nil, errPost
	}
	if postResponse.StatusCode == http.StatusOK {
		defer postResponse.Body.Close()
		body, err := io.ReadAll(postResponse.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &posts)
		if err != nil {
			return nil, err
		}

	}

	return posts, nil
}

func GetPost(postId string) (models.Post, error) {
	var post models.Post

	apiManager := manager.NewAPIManager()
	apiUrlManager := manager.NewAPIUrls()
	postApiUrl := apiUrlManager.GetPostsApiURL() + "/" + postId
	fmt.Println("Post API URL:", postApiUrl) // Log the post API URL

	postResponse, errPost := apiManager.Get(postApiUrl)
	if errPost != nil {

		return post, errPost
	}
	if postResponse.StatusCode == http.StatusOK {
		defer postResponse.Body.Close()
		body, err := io.ReadAll(postResponse.Body)
		if err != nil {
			return post, err
		}

		err = json.Unmarshal(body, &post)
		if err != nil {
			return post, err
		}

	}

	return post, nil
}

func AddPost(post models.Post) error {
	apiManager := manager.NewAPIManager()
	apiUrlManager := manager.NewAPIUrls()
	postApiUrl := apiUrlManager.GetPostsApiURL()
	fmt.Println("Post API URL:", postApiUrl) // Log the post API URL

	body, err := json.Marshal(post)
	if err != nil {
		return err
	}

	postResponse, errPost := apiManager.Post(postApiUrl, body)
	if errPost != nil {

		return errPost
	}
	if postResponse.StatusCode == http.StatusOK {
		defer postResponse.Body.Close()
		body, err := io.ReadAll(postResponse.Body)
		if err != nil {
			return err
		}

		fmt.Println("API Response Body:", string(body)) // Log the response body
	}
	fmt.Println("API Response Status:", postResponse.StatusCode) // Log the response status code
	return nil
}
