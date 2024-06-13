package manager

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

// APIManager handles API requests with proper authentication.
type APIManager struct {
	client     *http.Client
	guestToken string
	userToken  string
	isUserAuth bool
}

// NewAPIManager initializes the API manager with a guest token.
func NewAPIManager() *APIManager {
	return &APIManager{
		client:     &http.Client{Timeout: 10 * time.Second},
		guestToken: "0fc237962e95129004c313015d220aef4c7ffddc465cf984d1e63130b6e180c8",
		isUserAuth: false,
	}
}

// SetUserToken sets the user token and marks the user as authenticated.
func (a *APIManager) SetUserToken(token string) {
	a.userToken = token
	a.isUserAuth = true
}

// Request makes an API request with the appropriate authentication header.
func (a *APIManager) Request(method, url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set the Authorization header
	if a.isUserAuth {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.userToken))
	} else {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.guestToken))
	}

	req.Header.Set("Content-Type", "application/json")

	return a.client.Do(req)
}

// Get performs a GET request.
func (a *APIManager) Get(url string) (*http.Response, error) {
	return a.Request("GET", url, nil)
}

// Post performs a POST request.
func (a *APIManager) Post(url string, body []byte) (*http.Response, error) {
	return a.Request("POST", url, body)
}

// Put performs a PUT request.
func (a *APIManager) Put(url string, body []byte) (*http.Response, error) {
	return a.Request("PUT", url, body)
}

// Delete performs a DELETE request.
func (a *APIManager) Delete(url string) (*http.Response, error) {
	return a.Request("DELETE", url, nil)
}

func (a *APIManager) CisemFunc() {
	fmt.Println("Cisem")
}
