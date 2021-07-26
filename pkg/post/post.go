package post

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
