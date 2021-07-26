package handmadeMock

import "github.com/leoomi/GoTesting/pkg/post"

type AddPostCall struct {
	Post post.Post
}

type PostServiceHandmadeMock struct {
	GetPostCalled bool
	Posts         []post.Post

	AddPostCalled bool
	AddPostCalls  []AddPostCall
}

func (mock *PostServiceHandmadeMock) GetPosts() []post.Post {
	mock.GetPostCalled = true

	return mock.Posts
}

func (mock *PostServiceHandmadeMock) AddPost(post post.Post) {
	mock.AddPostCalled = true

	mock.AddPostCalls = append(mock.AddPostCalls,
		AddPostCall{
			Post: post,
		})
}
