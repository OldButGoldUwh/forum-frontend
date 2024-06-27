package manager

type apiUrls struct {
	GetUserApiUrl  string
	LoginApiUrl    string
	UserApiUrl     string
	PostsApiUrl    string
	PostApiUrl     string
	CommentsApiUrl string
}

const (
	GetUserApiUrl  = "http://localhost:8080/api/v1/user"
	LoginApiUrl    = "http://localhost:8080/api/v1/users/login"
	UserApiUrl     = "http://localhost:8080/api/v1/users"
	PostsApiUrl    = "http://localhost:8080/api/v1/posts"
	CommentsApiUrl = "http://localhost:8080/api/v1/comments"
)

func NewAPIUrls() *apiUrls {
	return &apiUrls{
		LoginApiUrl:    LoginApiUrl,
		GetUserApiUrl:  GetUserApiUrl,
		UserApiUrl:     UserApiUrl,
		PostsApiUrl:    PostsApiUrl,
		CommentsApiUrl: CommentsApiUrl,
	}
}

func (a *apiUrls) GetLoginApiURL() string {
	return a.LoginApiUrl
}

func (a *apiUrls) GetUserApiURL() string {
	return a.GetUserApiUrl
}

func (a *apiUrls) GetPostsApiURL() string {
	return a.PostsApiUrl
}

func (a *apiUrls) GetCommentsApiURL() string {
	return a.CommentsApiUrl
}
