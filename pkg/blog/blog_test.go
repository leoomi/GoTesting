package blog_test

import (
	"testing"

	"github.com/leoomi/GoTesting/pkg/blog"
	mock "github.com/leoomi/GoTesting/pkg/handmadeMock"
	"github.com/leoomi/GoTesting/pkg/post"
)

func TestGetBlog(t *testing.T) {
	tests := []struct {
		posts []post.Post
	}{
		{
			[]post.Post{{Id: 0}},
		},
		{
			[]post.Post{{Id: 1}},
		},
	}

	for _, test := range tests {
		mock := mock.PostServiceHandmadeMock{
			Posts: test.posts,
		}
		srv := blog.BlogServiceImpl{
			PostService: &mock,
		}

		blog := srv.GetBlog()

		if !mock.GetPostCalled {
			t.Error("PostService.GetPost not called")
		}

		if blog.Posts[0] != test.posts[0] {
			t.Error("Unexpected post returned")
		}
	}
}

func TestAddPost(t *testing.T) {
	tests := []struct {
		posts []post.Post
	}{
		{
			[]post.Post{{Id: 0}},
		},
		{
			[]post.Post{{Id: 1}, {Id: 2}},
		},
	}

	for _, test := range tests {
		mock := mock.PostServiceHandmadeMock{}
		srv := blog.BlogServiceImpl{
			PostService: &mock,
		}

		for _, post := range test.posts {
			srv.AddPost(post)
		}

		if !mock.AddPostCalled {
			t.Error("PostService.AddPost not called")
		}

		if len(test.posts) != len(mock.AddPostCalls) {
			t.Error("Unexpected amount of AddPost calls")
		}

		for i, post := range test.posts {
			if mock.AddPostCalls[i].Post != post {
				t.Error("Unexpected post used as parameter")
			}
		}
	}
}
