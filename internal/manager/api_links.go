package manager

type apiUrls struct {
	GetUserApiUrl     string
	LoginApiUrl       string
	UserApiUrl        string
	PostsApiUrl       string
	PostApiUrl        string
	CommentsApiUrl    string
	GetUsersApiUrl    string
	LikePostApiUrl    string
	DislikePostApiUrl string
}

const (
	GetUserApiUrl  = "http://localhost:8080/api/v1/user"
	GetUsersApiUrl = "http://localhost:8080/api/v1/users"

	LoginApiUrl       = "http://localhost:8080/api/v1/users/login"
	UserApiUrl        = "http://localhost:8080/api/v1/users"
	PostsApiUrl       = "http://localhost:8080/api/v1/posts"
	CommentsApiUrl    = "http://localhost:8080/api/v1/comments"
	LikePostApiUrl    = "http://localhost:8080/api/v1/posts/like/"
	DislikePostApiUrl = "http://localhost:8080/api/v1/posts/dislike/"
)

func NewAPIUrls() *apiUrls {
	return &apiUrls{
		LoginApiUrl:       LoginApiUrl,
		GetUserApiUrl:     GetUserApiUrl,
		GetUsersApiUrl:    GetUsersApiUrl,
		UserApiUrl:        UserApiUrl,
		PostsApiUrl:       PostsApiUrl,
		CommentsApiUrl:    CommentsApiUrl,
		LikePostApiUrl:    LikePostApiUrl,
		DislikePostApiUrl: DislikePostApiUrl,
	}
}

func (a *apiUrls) GetUsersApiURL() string {
	return a.GetUsersApiUrl
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

func (a *apiUrls) GetLikePostApiURL() string {
	return a.LikePostApiUrl
}

func (a *apiUrls) GetDislikePostApiURL() string {
	return a.DislikePostApiUrl
}
