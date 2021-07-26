package blog

import "github.com/leoomi/GoTesting/pkg/post"

type BlogServiceImpl struct {
	PostService post.PostService
}

var blog = Blog{
	Name: "Name",
}

func (srv *BlogServiceImpl) GetBlog() Blog {
	posts := srv.PostService.GetPosts()
	blog.Posts = posts

	return blog
}

func (srv *BlogServiceImpl) AddPost(post post.Post) {
	srv.PostService.AddPost(post)
}
