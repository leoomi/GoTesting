package blog_test

import (
	"testing"

	"github.com/leoomi/GoTesting/pkg/blog"
	mocks "github.com/leoomi/GoTesting/pkg/mockeryMocks"
	"github.com/leoomi/GoTesting/pkg/post"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBlogMockeryMock(t *testing.T) {
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
		postMock := mocks.PostService{}
		postMock.On("GetPosts").Return(test.posts)

		srv := blog.BlogServiceImpl{
			PostService: &postMock,
		}

		blog := srv.GetBlog()

		postMock.AssertCalled(t, "GetPosts")
		assert.Equal(t, blog.Posts[0], test.posts[0], "Returned post should be the same as mock")
	}
}

func TestAddPostMockeryMock(t *testing.T) {
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
		postMock := mocks.PostService{}
		postMock.On("AddPost", mock.AnythingOfTypeArgument("post.Post")).Return()

		srv := blog.BlogServiceImpl{
			PostService: &postMock,
		}

		for _, post := range test.posts {
			srv.AddPost(post)
		}

		postMock.AssertNumberOfCalls(t, "AddPost", len(test.posts))

		for _, post := range test.posts {
			postMock.AssertCalled(t, "AddPost", post)
		}
	}
}
