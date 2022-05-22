package blogpost_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"blogpost"
)

func assertPost(t *testing.T, got blogpost.Post, want blogpost.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1`
		secondBody = `Title: Post 2
Description: Description 2`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	// rest of test code cut for brevity
	posts, _ := blogpost.NewPostsFromFS(fs)
	assertPost(t, posts[0], blogpost.Post{
		Title:       "Post 1",
		Description: "Description 1",
	})
}
