package post

//go:generate mockgen -destination=../mockgenMock/PostService.go -package=mockgen_mocks . PostService
//go:generate mockery --name=PostService --output=../mockeryMocks
type Post struct {
	Id     int
	Title  string
	Body   string
	Author string // maybe new model
}

type PostService interface {
	GetPosts() []Post
	AddPost(post Post)
}
