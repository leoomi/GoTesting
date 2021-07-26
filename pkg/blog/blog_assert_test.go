package blog_test

import (
	"testing"

	"github.com/leoomi/GoTesting/pkg/blog"
	mock "github.com/leoomi/GoTesting/pkg/handmadeMock"
	"github.com/leoomi/GoTesting/pkg/post"

	"github.com/stretchr/testify/assert"
)

func TestGetBlogAssert(t *testing.T) {
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

		assert.False(t, mock.GetPostCalled, "GetPost method should be called")

		assert.Equal(t, blog.Posts[0], test.posts[0], "Returned post should be the same as mock")
	}
}

func TestAddPostAssert(t *testing.T) {
	tests := []struct {
		Posts []post.Post
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

		for _, post := range test.Posts {
			srv.AddPost(post)
		}

		assert.True(t, mock.AddPostCalled, "PostService.AddPost not called")

		assert.Equal(t, len(test.Posts), len(mock.AddPostCalls), "Unexpected amount of AddPost calls")

		for i, post := range test.Posts {
			assert.Equal(t, mock.AddPostCalls[i].Post, post, "Unexpected post used as parameter")
		}
	}
}
