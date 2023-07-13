package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/zypherscript/blogposts"
)

// func TestNewBlogPostsOld(t *testing.T) {
// 	fs := fstest.MapFS{
// 		"hello world.md":  {Data: []byte("hi")},
// 		"hello-world2.md": {Data: []byte("hola")},
// 	}

// 	posts, err := blogposts.NewPostsFromFS(fs)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(posts) != len(fs) {
// 		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
// 	}
// }

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogPostsFail(t *testing.T) {
	_, err := blogposts.NewPostsFromFS(StubFailingFS{})

	if err == nil {
		t.Error("expected an error but didn't get one")
	}
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title: Post 2")},
	}

	posts, _ := blogposts.NewPostsFromFS(fs)
	assertPost(t, posts[0], blogposts.Post{Title: "Post 1"})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
