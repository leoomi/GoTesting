package post

type PostServiceImpl struct {
}

var posts = []Post{
	{
		Id:     0,
		Title:  "Title 0",
		Body:   "Body 0",
		Author: "Author 0",
	},
	{
		Id:     1,
		Title:  "Title 1",
		Body:   "Body 1",
		Author: "Author 1",
	},
}

func (srv *PostServiceImpl) GetPosts() []Post {
	return posts
}

func (srv *PostServiceImpl) AddPost(post Post) {
	posts = append(posts, post)
}
