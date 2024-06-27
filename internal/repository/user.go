package repository

import (
	"encoding/json"
	"fmt"
	"golang-forum-frontend/internal/manager"
	"golang-forum-frontend/internal/models"
	"io"
	"net/http"
)

func GetUsername() (string, error) {
	fmt.Println("Get Username")
	var username string

	apiManager := manager.NewAPIManager()

	apiUrlManager := manager.NewAPIUrls()
	userApiUrl := apiUrlManager.GetUserApiURL()

	userResponse, err := apiManager.Get(userApiUrl)

	if err != nil {
		fmt.Println("Error:", err)
		return username, err
	}

	if userResponse.StatusCode == http.StatusOK {
		defer userResponse.Body.Close()
		body, err := io.ReadAll(userResponse.Body)
		fmt.Println("Body:", string(body))

		if err != nil {
			return username, err
		}

		var user models.User
		fmt.Println("User:", user)

		err = json.Unmarshal(body, &user)
		if err != nil {
			return username, err
		}

		fmt.Println("User:", user)

		username = user.Username
	}

	return username, nil

}
