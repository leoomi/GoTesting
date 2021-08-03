package blog_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/leoomi/GoTesting/pkg/blog"
	mocks "github.com/leoomi/GoTesting/pkg/mockgenMock"
	"github.com/leoomi/GoTesting/pkg/post"

	"github.com/stretchr/testify/assert"
)

func TestGetBlogGoMock(t *testing.T) {
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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range tests {
		postMock := mocks.NewMockPostService(ctrl)
		postMock.
			EXPECT().
			GetPosts().
			Return(test.posts).
			Times(1)

		srv := blog.BlogServiceImpl{
			PostService: postMock,
		}

		blog := srv.GetBlog()

		assert.Equal(t, blog.Posts[0], test.posts[0], "Returned post should be the same as mock")
	}
}

func TestAddPostGoMock(t *testing.T) {
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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range tests {
		postMock := mocks.NewMockPostService(ctrl)

		// Two expects will make one of them fail. Can't replicate what was done in the other tests
		// postMock.
		// 	EXPECT().
		// 	AddPost(gomock.AssignableToTypeOf(post.Post{})).
		// 	Times(len(test.posts))

		for _, post := range test.posts {
			postMock.
				EXPECT().
				AddPost(gomock.Eq(post)).
				Times(1)
		}

		srv := blog.BlogServiceImpl{
			PostService: postMock,
		}

		for _, post := range test.posts {
			srv.AddPost(post)
		}
	}
}

func TestAddPostInOrderGoMock(t *testing.T) {
	test := struct {
		posts []post.Post
	}{
		[]post.Post{{Id: 0}, {Id: 1}},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postMock := mocks.NewMockPostService(ctrl)

	gomock.InOrder(
		postMock.EXPECT().AddPost(test.posts[0]),
		postMock.EXPECT().AddPost(test.posts[1]),
	)

	srv := blog.BlogServiceImpl{
		PostService: postMock,
	}

	for _, post := range test.posts {
		srv.AddPost(post)
	}
}
