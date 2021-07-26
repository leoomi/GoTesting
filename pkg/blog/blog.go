package blog

import "github.com/leoomi/GoTesting/pkg/post"

type Blog struct {
	Name  string
	Posts []post.Post
}

type BlogService interface {
	GetBlog() Blog
}
